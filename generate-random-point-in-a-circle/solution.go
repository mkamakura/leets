type Solution struct {
    r, xc, yc float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
    return Solution{radius, x_center, y_center}
}

func (this *Solution) RandPoint() []float64 {
    x0 := this.xc - this.r
    y0 := this.yc - this.r

    for true {
        xg := x0 + rand.Float64() * 2 * this.r
        yg := y0 + rand.Float64() * 2 * this.r
        if math.Sqrt((xg-this.xc)*(xg-this.xc) + (yg-this.yc)*(yg-this.yc)) <= this.r {
            return []float64{xg, yg}
        }
    }

    return []float64{}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
