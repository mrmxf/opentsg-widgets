# tpg-widgets

`opentsg-widgets` is the library of builtin widgets.
Feel free to use any as a demo to make your own widgets.

## Examples

Examples of the json for each type can be found at the `exampleJson` folder.
These examples contain all the fields unique to the widget. The name of the folder
is the same as the type that would be declared for that widget. examples that have nested
folders such as builtin.ebu3373/bars, have their widget names including the `/`, so
`builtin.ebu3373/bars` would be the widget type.

The positional fields of grid are expected to be in **every** widget. And have the
layout as shown below. They may not be found on every example json.

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

The following  sections includes a description of the widgets and their
intended design use and fields and variables.

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

EBU 3373 is a test pattern for testing HD Hybrid Log Gamma workflows,
for OpenTSG the individual segments of the test pattern have been separated
into widgets , so that you can use just the components you require.
More technical details of each segment
 can be found [here.](https://tech.ebu.ch/publications/tech3373)

#### Bars

The ebu3373 bars are generated to fill the space.
The bars are in 10 bit BT2020 colour space.

It has no fields, only the location is required.

#### Luma

A horizontal luma ramp from the 10 bit value 4 to 1019,
 is generated to fill the widget area.

It has no fields, only the location is required.

#### Near black

Alternating segments of 10 bit 0% black (RGB10(64,64,64)) and
-4%, -2%, -1%, 1%, 2%, 4% black. This is for checking
sub blacks are not removed in the production chain.
 This widget is generated to fill the widget area.

It has no fields, only the location is required.

#### Saturation

The saturation generates the steps of red, green and blue
from minimum to maximum saturation. These saturations
are used to establish the colour space of the viewing device.

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

TWO Si is the two sample interleave pattern, for when SMPTE ST 425-5:2019
is used. The order of the four 3G-SDI video cables can be checked.
It consists of 8 combos of the four ABCD cables and lines. Each letter and 
set of lines are linked to a cable.

It has no fields, only the location is required.

### four Colour

Is the four colour algorithm for TPIGS.Where four
colours are the maximum number of colours to
fill a map. This widget allows 4+ colours to be used
in the interest of computational speed for large objects.

- `colors` - a list of TSG colors. The algortihim will use the least amount of colors required.

### framecount

Produces a four number long frame number, for
where the test card is in a pattern.

### Noise

Noise generates 12 bit white noise to fill the widget.
It has the following fields

- `minimum` - the minimum 12 bit integer value for the white noise
- `maximum` - the maximum 12 bit integer value for the white noise

```json
{
    "noisetype": "white noise",
    "minimum": 0,
    "maximum": 4095
}
```

### qr gen

Generates a qr code from the user input.
It has the following fields

```json
    "code": "https://opentsg.io/",
    "gridPosition": {
        "x": 97.1,
        "y": 97.1
    }
```

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

Generates a zoneplate to fill the widget.