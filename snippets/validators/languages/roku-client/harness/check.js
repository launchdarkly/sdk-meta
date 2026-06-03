// Parse the staged .brs file via brighterscript's Lexer + Parser.
// Exit 0 on clean parse; print diagnostics to stderr and exit 1
// otherwise.
//
// brighterscript ships its parser as a library on the global
// install path; resolve via require.resolve so we pick up the
// globally-installed copy from the Dockerfile's `npm install -g`.
const path = require('path');
const fs = require('fs');

// The Dockerfile sets NODE_PATH=/usr/local/lib/node_modules so the
// globally-installed brighterscript resolves through normal require().
const bs = require('brighterscript');

const file = process.argv[2];
if (!file) {
    console.error('usage: node check.js <file.brs>');
    process.exit(2);
}

const src = fs.readFileSync(file, 'utf8');

// Lex.
const lex = bs.Lexer.scan(src);
const lexErrors = (lex.diagnostics || []).filter(
    (d) => (d.severity || 1) === 1, // 1 = Error in brighterscript's DiagnosticSeverity enum
);

// Parse.
const parse = bs.Parser.parse(lex.tokens || []);
const parseErrors = (parse.diagnostics || []).filter(
    (d) => (d.severity || 1) === 1,
);

const errors = [...lexErrors, ...parseErrors];
if (errors.length > 0) {
    for (const d of errors) {
        const line =
            d.range && d.range.start ? d.range.start.line + 1 : '?';
        const col =
            d.range && d.range.start
                ? d.range.start.character + 1
                : '?';
        console.error(`${path.basename(file)}:${line}:${col}: ${d.message}`);
    }
    process.exit(1);
}
