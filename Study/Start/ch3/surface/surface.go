package surface

import (
    "fmt"
    "io"
    "math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)


func DrawSurfaceWithParams(out io.Writer, width, height int, color string) {
    // 使用传入的宽度和高度
    const cells = 100
    xyrange := 30.0
    xyscale := float64(width) / 2 / xyrange
    zscale := float64(height) * 0.4
    angle := math.Pi / 6
    sin30, cos30 := math.Sin(angle), math.Cos(angle)

    fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay := corner(i+1, j, xyscale, zscale, sin30, cos30, xyrange, cells)
            bx, by := corner(i, j, xyscale, zscale, sin30, cos30, xyrange, cells)
            cx, cy := corner(i, j+1, xyscale, zscale, sin30, cos30, xyrange, cells)
            dx, dy := corner(i+1, j+1, xyscale, zscale, sin30, cos30, xyrange, cells)

            if isInvalid(ax, ay) || isInvalid(bx, by) || isInvalid(cx, cy) || isInvalid(dx, dy) {
                continue
            }

            // Calculate average height for the polygon
            z1 := f(xyrange*(float64(i+1)/cells-0.5), xyrange*(float64(j)/cells-0.5))
            z2 := f(xyrange*(float64(i)/cells-0.5), xyrange*(float64(j)/cells-0.5))
            z3 := f(xyrange*(float64(i)/cells-0.5), xyrange*(float64(j+1)/cells-0.5))
            z4 := f(xyrange*(float64(i+1)/cells-0.5), xyrange*(float64(j+1)/cells-0.5))
            avgZ := (z1 + z2 + z3 + z4) / 4

            // Map height to color
            fillColor := heightToColor(avgZ)
            if color != "" {
                fillColor = color // 使用客户端指定的颜色
            }

            fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy, fillColor)
        }
    }
    fmt.Println("</svg>")
}

func corner(i, j int, xyscale, zscale, sin30, cos30, xyrange float64, cells int) (float64, float64) {
    // Find point (x,y) at corner of cell (i,j).
    x := xyrange * (float64(i)/float64(cells) - 0.5)
    y := xyrange * (float64(j)/float64(cells) - 0.5)

    // Compute surface height z.
    z := f(x, y)

    // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
    sx := float64(width)/2 + (x-y)*cos30*xyscale
    sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y) // 距离原点的半径
    return math.Sin(r) / 10
}

func isInvalid(x, y float64) bool {
    return math.IsNaN(x) || math.IsNaN(y) || math.IsInf(x, 0) || math.IsInf(y, 0)
}

func heightToColor(z float64) string {
    // Normalize z to a range of 0 to 1
    minZ, maxZ := -0.1, 0.1 // Adjust based on the expected range of z
    normalized := (z - minZ) / (maxZ - minZ)
    if normalized < 0 {
        normalized = 0
    } else if normalized > 1 {
        normalized = 1
    }

    // Interpolate between blue and red
    red := int(255 * normalized)
    blue := int(255 * (1 - normalized))
    return fmt.Sprintf("#%02x00%02x", red, blue)
}