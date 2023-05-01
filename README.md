A schematics engine written in go in-order to innovate the schematics niche and increase development speed.

### Install
`npm install -D hopemanryan/gschema`


### How to use

#### Templates 

 * Templates files must end with a  `__templ__` suffix 
 * The regex pattern for dynamic content is: `<% file_name =%>` (in the context file_name is a variable)
 * You can add dynamic variables to file names `<% file_name =%>Component.ts__templ__`
 * Javascript functions are supported (see more down bellow)

 #### Example

 ```
 const <% toUpperCase(file_name) =%>Component = () => {

}


const <% file_name =%>Component = () => {

}
 ```


#### Basic usage

```bash 
node ./node_modules/gschema/gschema.js -file_name=demo1 -read_dir=./abc
```


### JS functions 

Schematics are more popular in the Javascript community. There for we have added support to add a JS file that will have functions which can be used in the template files 

```javascript
// gschema.js
function toUpperCase (val) {
    return val.toUpperCase();
}
```

This will load the function and the engine sees toUpperCase it will run the function with variable value as the argument  `toUpperCase('demo1')` and return the value `DEMO1`


### Shorthands

`gschema.config.json` is the config file that can be used for shorthands configuration

```json
{
    "shorthands": {
        "demo": {
            "templatePath": "./abc"
        }
    }
}
```
``` bash
node ./node_modules/gschema/gschema -file_name=demo1 -s=demo
```
The engine will use the predefined config, this allows true re-usability of templates