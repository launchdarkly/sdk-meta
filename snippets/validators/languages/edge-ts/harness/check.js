// Parse-only TypeScript check via the compiler's transpileModule.
// transpileModule reports only SYNTACTIC diagnostics -- it does not
// resolve modules or type-check -- so edge fragments importing
// uninstalled packages and referencing ambient globals (env, process)
// pass as long as they parse. Exit 0 on clean parse; print diagnostics
// to stderr and exit 1 otherwise.
const fs = require('fs');
const path = require('path');
// typescript is installed under /opt/edge-ts by the Dockerfile; the
// harness runs from /work, so resolve it by absolute path.
const ts = require('/opt/edge-ts/node_modules/typescript');

const file = process.argv[2];
if (!file) {
    console.error('usage: node check.js <file.ts>');
    process.exit(2);
}

const src = fs.readFileSync(file, 'utf8');
const out = ts.transpileModule(src, {
    compilerOptions: {
        target: ts.ScriptTarget.ES2020,
        module: ts.ModuleKind.ESNext,
        jsx: ts.JsxEmit.Preserve,
    },
    reportDiagnostics: true,
    fileName: path.basename(file),
});

const errors = (out.diagnostics || []).filter(
    (d) => d.category === ts.DiagnosticCategory.Error,
);
if (errors.length > 0) {
    for (const d of errors) {
        const msg = ts.flattenDiagnosticMessageText(d.messageText, '\n');
        let loc = '?';
        if (d.file && d.start != null) {
            const p = d.file.getLineAndCharacterOfPosition(d.start);
            loc = `${p.line + 1}:${p.character + 1}`;
        }
        console.error(`${path.basename(file)}:${loc}: ${msg}`);
    }
    process.exit(1);
}
