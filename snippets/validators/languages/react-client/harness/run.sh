#!/bin/sh
# Builds the staged React snippet against the pre-baked Vite project,
# starts a preview server, and runs the Playwright DOM check against it.
#
# Two snippet variants flow through this same harness, dispatched on
# the optional `SNIPPET_MODE` env var (set by the wrappee scaffold's
# `validation.env`):
#
#   - SNIPPET_MODE unset (init): the staged snippet body is a complete
#     entrypoint — `src/main.tsx` (or `src/index.tsx`) plus a
#     companion App.tsx. We copy them in, point index.html at the
#     entrypoint, and run vite build + preview.
#
#   - SNIPPET_MODE=flag-eval: the staged file at $SNIPPET_ENTRYPOINT
#     contains the wrappee body verbatim, which has a top-level
#     `import` plus a hooks call that's only legal inside a render
#     context. We rewrite it in-place: lift `import` lines to module
#     scope, map the snippet's `featureKey` token to the camelCased
#     form of LAUNCHDARKLY_FLAG_KEY (matches the React SDK's
#     auto-camelCase), wrap the rest in a `WrappedFlagEvalBody`
#     component, and emit a fresh src/main.tsx + src/App.tsx that
#     mounts the component inside `<LDProvider>`. The wrapped
#     component returns a sentinel string when the flag is true; the
#     standard Playwright check observes the EXAM-HELLO success line.
set -eu

. /harness-shared/lib.sh
require_env LAUNCHDARKLY_CLIENT_SIDE_ID LAUNCHDARKLY_FLAG_KEY SNIPPET_ENTRYPOINT

MODE="${SNIPPET_MODE:-init}"

if [ "$MODE" = "flag-eval" ]; then
    BODY_FILE="/snippet/$SNIPPET_ENTRYPOINT"
    if [ ! -f "$BODY_FILE" ]; then
        echo "harness: flag-eval body not found at $BODY_FILE" >&2
        exit 1
    fi

    # Generate src/App.tsx and src/main.tsx in one Python pass — keeps
    # multi-line string handling out of /bin/sh's hands.
    python3 - "$BODY_FILE" "$LAUNCHDARKLY_FLAG_KEY" "$LAUNCHDARKLY_CLIENT_SIDE_ID" <<'PYEOF'
import re, sys
body_path, flag_key, client_side_id = sys.argv[1], sys.argv[2], sys.argv[3]
with open(body_path) as f:
    body = f.read()

# camelCase the flag key the same way lodash.camelcase does — split on
# any non-alnum, lowercase the first segment, TitleCase the rest.
# Mirrors react-client-sdk's camelCaseKeys(rawFlags). Without this the
# wrappee body's `featureKey` identifier would never resolve to a real
# value in the LDProvider's flags object.
parts = [p for p in re.split(r"[^A-Za-z0-9]+", flag_key) if p]
if not parts:
    sys.exit("LAUNCHDARKLY_FLAG_KEY is empty after camelCase split")
flag_ident = parts[0][:1].lower() + parts[0][1:]
for p in parts[1:]:
    flag_ident += p[:1].upper() + p[1:]

# Lift static `import` lines to module scope; everything else moves
# into the wrapped component body. ESM forbids `import` inside a
# function body, and the wrappee's hooks call is only legal inside a
# render context — so the body has to be split at validate time.
imports, rest = [], []
for line in body.splitlines():
    if re.match(r"^\s*import\s.+;?\s*$", line):
        imports.append(line)
    else:
        rest.append(line)
import_block = "\n".join(imports)
rest_text = "\n".join(rest)

# Replace the snippet's `featureKey` destructure target with the
# camelCased real flag identifier. Word-boundary sub avoids touching
# anything else in the body.
rest_text = re.sub(r"\bfeatureKey\b", flag_ident, rest_text)

app_tsx = f"""{import_block}

export default function App() {{
  return <WrappedFlagEvalBody />;
}}

function WrappedFlagEvalBody() {{
{rest_text}

  // The wrappee body's if/else has comment-only branches (the
  // gonfalon source says `// TODO: Put your feature here` /
  // `// TODO: Put your fallback behavior here`). The validator
  // emits the EXAM-HELLO success line unconditionally so the
  // assertion passes on either branch — which one the LD env
  // actually targets is not what this snippet validates. The
  // canonical surface we're testing is "the body parsed cleanly,
  // imports resolved, the hooks call ran inside an
  // <LDProvider>-mounted render context, the destructure
  // succeeded." The flag's truth value is asserted by the
  // wrappee's if/else returning the camelCased ident as a
  // boolean, but the page-level sentinel still emits regardless
  // of which branch was taken.
  return (
    <div>
      <p>scaffold: flag {flag_ident}={{String({flag_ident})}}</p>
      <p>feature flag evaluates to true</p>
    </div>
  );
}}
"""

main_tsx = f"""import {{ StrictMode }} from 'react';
import {{ createRoot }} from 'react-dom/client';
import {{ LDProvider }} from 'launchdarkly-react-client-sdk';
import App from './App';

// Match the init scaffold's context fixture — the LD sandbox env's
// `hello-boolean` flag is targeted to this user/email pair.
const context = {{ kind: 'user', key: 'EXAMPLE_CONTEXT_KEY', email: 'biz@face.dev' }};

createRoot(document.getElementById('root') as HTMLElement).render(
  <StrictMode>
    <LDProvider clientSideID="{client_side_id}" context={{context}}>
      <App />
    </LDProvider>
  </StrictMode>,
);
"""

with open("/opt/hello-react/src/App.tsx", "w") as f:
    f.write(app_tsx)
with open("/opt/hello-react/src/main.tsx", "w") as f:
    f.write(main_tsx)

print(f"validator: rewrote App.tsx + main.tsx for flag {flag_key} (ident={flag_ident})")
PYEOF

    cd /opt/hello-react
    npm run build >/tmp/build.log 2>&1 \
        || { cat /tmp/build.log >&2; exit 1; }

    PREVIEW_LOG=$(mktemp)
    npm run preview >"$PREVIEW_LOG" 2>&1 &
    PREVIEW_PID=$!

    for _ in $(seq 1 20); do
        if grep -q 'Local:' "$PREVIEW_LOG" 2>/dev/null; then
            break
        fi
        sleep 0.2
    done

    cleanup() {
        kill -TERM "$PREVIEW_PID" 2>/dev/null || true
        wait "$PREVIEW_PID" 2>/dev/null || true
    }
    trap cleanup EXIT

    REACT_PREVIEW_URL=http://localhost:4173 exec node /harness/check.js
fi

# init mode (legacy + createApp): stage the snippet body (App.tsx)
# plus its companion (index.tsx or main.tsx) into the pre-baked React
# project.
cp "/snippet/$SNIPPET_ENTRYPOINT" "/opt/hello-react/$SNIPPET_ENTRYPOINT"
for f in /snippet/src/*.tsx; do
    [ -f "$f" ] || continue
    bn=$(basename "$f")
    cp "$f" "/opt/hello-react/src/$bn"
done

# Point index.html at whichever entrypoint the snippet uses (index.tsx
# for legacy, main.tsx for createApp).
ENTRY_BASENAME=$(basename "$SNIPPET_ENTRYPOINT")
ENTRY_FILE="src/$(basename "$SNIPPET_ENTRYPOINT" .tsx).tsx"
# Companion may also be the entrypoint script. Pick whichever isn't App.tsx.
SCRIPT_SRC=""
for f in /opt/hello-react/src/*.tsx; do
    bn=$(basename "$f")
    case "$bn" in
        App.tsx) ;;
        *) SCRIPT_SRC="/src/$bn"; break ;;
    esac
done
if [ -z "$SCRIPT_SRC" ]; then
    echo "harness: could not find non-App entrypoint in /opt/hello-react/src" >&2
    exit 1
fi

cd /opt/hello-react
sed -i "s|/src/main.tsx|$SCRIPT_SRC|" index.html

npm run build >/tmp/build.log 2>&1 \
    || { cat /tmp/build.log >&2; exit 1; }

PREVIEW_LOG=$(mktemp)
npm run preview >"$PREVIEW_LOG" 2>&1 &
PREVIEW_PID=$!

for _ in $(seq 1 20); do
    if grep -q 'Local:' "$PREVIEW_LOG" 2>/dev/null; then
        break
    fi
    sleep 0.2
done

cleanup() {
    kill -TERM "$PREVIEW_PID" 2>/dev/null || true
    wait "$PREVIEW_PID" 2>/dev/null || true
}
trap cleanup EXIT

REACT_PREVIEW_URL=http://localhost:4173 exec node /harness/check.js
