var fs = require('fs');
var CodeGen = require('swagger-js-codegen').CodeGen;

var file = 'my_service.swagger.json';
var swagger = JSON.parse(fs.readFileSync(file, 'UTF-8'));
var reactjsSourceCode = CodeGen.getReactCode({ className: 'MyService', swagger: swagger });
console.log(reactjsSourceCode);
