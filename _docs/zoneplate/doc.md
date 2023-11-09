# Zoneplate

Generates a zoneplate to fill the given area.

It has the following fields:

- `platetype` - the type and shape of zone plate these are circular, sweep or ellipse.
The default option is circular.
- `startcolor` -  the start colour of the zone plate, these are white, black, grey or gray.
- `angle` - the angle for the zone plate to be to rotated, fits the following patterns
`^π\\*(\\d){1,4}$`, `^π\\*(\\d){1,4}/{1}(\\d){1,4}$` or a number between 0 and 360

```json
{
    "platetype": "sweep",
    "startcolor": "white",
    "angle": "π*1",
}
```
