const { exec } = require('child_process');
const path = require('path')
const fullPath = path.join(__dirname, '/initial-service')
const goArgs = process.argv.slice(2).join(' ')
exec(` ${fullPath} ${goArgs}`,
        (error, stdout, stderr) => {
            console.log(stdout);
            console.log(stderr);
            if (error !== null) {
                console.log(`exec error: ${error}`);
            }
});
