module.exports = {
    entry: './build/frontend/main.js', // tsc outputs to this directory.
    output: {
        filename: './build/main.js'  
    },
    externals: {
        'react': 'React',
        'react-dom': 'ReactDOM'
    }
}