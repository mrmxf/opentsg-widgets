# tpg-widgets

`opentsg-widgets` is the library of builtins. Documentation to follow.

## Examples

Examples of the json for each type can be found at the `exampleJson` folder.
These examples contain all the fields unique to the widget. The name of the folder
is the same as the type that would be declared for that widget. examples that have nested
folders such as builtin.ebu3373/bars, have their widget names including the `/`, so
`builtin.ebu3373/bars` would be the widget type.

The positional fields of grid are expected to be in **every** widget. And have the
layout as shown below. They may not be found on every example.



```json
"type" : "builtin.example",
"grid": {
    "location": "a1:b2",
    "alias" : "A demo Alias"
}

```
