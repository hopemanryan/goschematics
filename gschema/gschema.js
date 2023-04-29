const { exec } = require('child_process');
const path = require('path')
const fullPath = path.join(process.cwd(), 'node_modules','gschema','initial-service')
console.log(fullPath)
const goArgs = process.argv.slice(2).join(' ') + ` -_entrypoint=${process.cwd()}`

exec(` ${fullPath} ${goArgs}`,
        (error, stdout, stderr) => {
            console.log(stdout);
            console.log(stderr);
            if (error !== null) {
                console.log(`exec error: ${error}`);
            }
});
