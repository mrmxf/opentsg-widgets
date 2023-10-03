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

## Widget Properties

All widget have the Grid layout field, saying where the widget
is to be located on the test pattern.

General layout 

- what the widget sets out to do
- fields / talk about the layout
- Demo Json
- quirks Widget usage

The following includes a description of the widgets and their intended design use and fields and variables.

### Add Image

Add image adds an image onto the test pattern.

It has the following fields:

- `image` - the local or online location of image to be added.

- `imageFill` - the fill type of the image
  - `fill` - Stretch the image to fill the X and Y axis
  - `x scale` - Scale the image to fit the X axis
  - `y scale` - Scale the image to fit the Y axis
  - `xy scale` -  Scale the image to fit the X or Y axis, which ever one is reached first.

```json
{
    "image": "./exmaple/exampleImage.png",
    "imageFill": "fill"
}
```

### EBU3733

#### Bars

The ebu3373 bars are generated to fill the space.

It has no fields, only the location is required.

#### Luma

A horizontal luma ramp from x to y is generated to fill the widget area.

It has no fields, only the location is required.

#### Near black

A horizontal Near black ramo from x y to is generated to fill the widget area.

It has no fields, only the location is required.

#### Saturation

The saturation generates the RGB saturation steps

It has the following fields

- `colors` - An array of the "red", "green" and "blue". The saturation appears in the order
they were declared. If only one colour is chosen then only that colour is used.

```json
{
    "colors": [
        "red",
        "green",
        "blue"
    ]
}
```

#### Two SI

TWO Si is the two sample interleave pattern, for when four cables are used. Can check the ABCD order.
The pattern is Four sets of letters of ABCD

It has no fields, only the location is required.

### four Colour

Is the four colour algorthim for TPIGS. Optional up to four plus colours. Runtimes will decrease with more colours

- `colors` - a list of TSG colors. The algortihim will use the least amount of colors required.

### framecount 

Declare what frame we is in a frane order

### qr gen

Generates a qr code.

### Gradients

Gradients are for checking bit depths of a display, up to 16 bits.
Each gradient can be set up to shift a certain bit depths, the
shift is in relation to the bitdepth, if two alternate bit depths appear
to have the same shift then the limit of the display may have been found.

Gradients has the following fields

- Gradients
- Groups
- WidgetProperties

### Textbox

generates a text box of one or more lines. Lots of fonts etc

### Zoneplate

Generates a zoneplate to fill