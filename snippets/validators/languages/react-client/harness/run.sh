#!/bin/sh
# Builds the staged React snippet against the pre-baked Vite project,
# starts a preview server, and runs the Playwright DOM check against it.
#
# Two snippet variants flow through this same harness, dispatched on
# the optional `SNIPPET_MODE` env var (set by the wrappee scaffold's
# `validation.env`):
#
#   - SNIPPET_MODE unset (init): the staged snippet body is a complete
#     entrypoint -- `src/main.tsx` plus a companion App.tsx. We copy
#     them in and run vite build + preview.
#
#   - SNIPPET_MODE=flag-eval: the staged file at $SNIPPET_ENTRYPOINT
#     contains the wrappee body verbatim, which has a top-level
#     `import` plus a hooks call that's only legal inside a render
#     context. We rewrite it in-place: lift `import` lines to module
#     scope, wrap the rest in a `WrappedFlagEvalBody` component, and
#     emit a fresh src/main.tsx + src/App.tsx that mounts the
#     component inside `<LDReactProvider>`. The wrapping component
#     unconditionally renders the EXAM-HELLO success line; the
#     standard Playwright check observes it.
#
#     The wrappee's flag-key string is substituted with the live
#     LAUNCHDARKLY_FLAG_KEY upstream via the snippet's
#     `validation.placeholders` block, so this harness does not need
#     to do any flag-key rewriting.
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

    # Generate src/App.tsx and src/main.tsx in one Python pass -- keeps
    # multi-line string handling out of /bin/sh's hands.
    python3 - "$BODY_FILE" "$LAUNCHDARKLY_CLIENT_SIDE_ID" <<'PYEOF'
import re, sys
body_path, client_side_id = sys.argv[1], sys.argv[2]
with open(body_path) as f:
    body = f.read()

# Lift static `import` lines to module scope; everything else moves
# into the wrapped component body. ESM forbids `import` inside a
# function body, and the wrappee's hooks call is only legal inside a
# render context -- so the body has to be split at validate time.
imports, rest = [], []
for line in body.splitlines():
    if re.match(r"^\s*import\s.+;?\s*$", line):
        imports.append(line)
    else:
        rest.append(line)
import_block = "\n".join(imports)
rest_text = "\n".join(rest)

app_tsx = f"""{import_block}

export default function App() {{
  return <WrappedFlagEvalBody />;
}}

function WrappedFlagEvalBody() {{
{rest_text}

  // The wrappee's if/else has comment-only branches (`// TODO: Put
  // your feature here` / `// TODO: Put your fallback behavior here`).
  // The validator emits the EXAM-HELLO success line unconditionally
  // so the assertion passes regardless of which branch executes --
  // we are testing that the body parsed, imports resolved, and the
  // hook ran inside an <LDReactProvider>-mounted render context, not
  // the flag's truth value.
  return <p>feature flag evaluates to true</p>;
}}
"""

main_tsx = f"""import {{ createRoot }} from 'react-dom/client';
import {{ createLDReactProvider, LDContext }} from '@launchdarkly/react-sdk';
import App from './App';

// Match the init scaffold's context fixture -- the LD sandbox env's
// boolean flag is targeted to this user/email pair.
const context: LDContext = {{ kind: 'user', key: 'EXAMPLE_CONTEXT_KEY', email: 'biz@face.dev' }};
const LDReactProvider = createLDReactProvider('{client_side_id}', context);

createRoot(document.getElementById('root') as HTMLElement).render(
  <LDReactProvider>
    <App />
  </LDReactProvider>,
);
"""

with open("/opt/hello-react/src/App.tsx", "w") as f:
    f.write(app_tsx)
with open("/opt/hello-react/src/main.tsx", "w") as f:
    f.write(main_tsx)

print("validator: rewrote App.tsx + main.tsx for flag-eval body")
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
