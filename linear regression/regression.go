// * using three packages for plotting "gonum.org/v1/plot","gonum.org/v1/plot/plotter","gonum.org/v1/plot/vg"
// * i am considering an car dataset for linear regression

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Loading the dataset
	file, err := os.Open("C:/Users/Vissarapu Srinath/Downloads/car data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Convert the data to floats
	var x []float64
	var y []float64
	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}
		xVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		yVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		x = append(x, xVal)
		y = append(y, yVal)
	}

	// here i am computing the coefficients into linear regression lines
	var sumX, sumY, sumXY, sumX2 float64
	n := float64(len(x))
	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += math.Pow(x[i], 2)
	}
	m := (n*sumXY - sumX*sumY) / (n*sumX2 - math.Pow(sumX, 2))
	b := (sumY - m*sumX) / n

	// by this formula i am writing an liner equation
	fmt.Printf("y = %.2fx + %.2f\n", m, b)

	// Plotting the data points and regression line
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	points := make(plotter.XYs, len(x))
	for i := range points {
		points[i].X = x[i]
		points[i].Y = y[i]
	}

	line := plotter.NewFunction(func(x float64) float64 { return m*x + b })

	p.Add(plotter.NewGrid())
	p.Add(plotter.NewScatter(points))
	p.Add(line)

	// Setting plot labels
	p.Title.Text = "Linear Regression"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Save the plot to a PNG file
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "regression.png"); err != nil {
		log.Fatal(err)
	}
}
