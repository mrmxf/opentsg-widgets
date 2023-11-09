# framecount

Produces a four number long frame number, for
where the test card is in a sequence of patterns.

It has the following properties.

- `font` - the font to be used, it can be an in built font of title, body,
pixel or header. Or it can be the path to a local or web file.
- `framecounter` - if true then the frame counter will be used. If false then
the widget is skipped.
- `textcolor` - the colour of the text.
- `backgroundcolor` - the colour of the background.
- `fontsize` - the font size of the frame counter,
this dictates the size of the frame counter box
- `gridPosition` - the relative x,y positions as percentages
of the grid the inhabit. There are also the builtin positions of
`"bottom right"`, `"bottom left"`,`"top right"` or `"top left"`.

```json
    "framecounter": true,
    "textcolor": "",
    "backgroundcolor": "",
    "font": "",
    "fontsize": 22, 
    "gridPosition": "top left"
```
