package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type BenchmarkResult struct {
	Name        string
	Depth       int
	Iterations  int
	TimePerOp   int
	BytesPerOp  int
	AllocsPerOp int
}

func main() {
	file, err := os.Open("benchmark_results.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	var results []BenchmarkResult
	scanner := bufio.NewScanner(file)

	// Refined regex to capture benchmark results accurately
	re := regexp.MustCompile(`^BenchmarkInvestigateTreeGeneral/([\w/]+)(-d\d+)-\w+\s+(\d+)\s+(\d+)\s+ns/op\s+(\d+)\s+B/op\s+(\d+)\s+allocs/op`)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		if match != nil {
			name := match[1]
			depth, _ := strconv.Atoi(strings.Trim("-d", match[2]))
			iterations, _ := strconv.Atoi(match[3])
			timePerOp, _ := strconv.Atoi(match[4])
			bytesPerOp, _ := strconv.Atoi(match[5])
			allocsPerOp, _ := strconv.Atoi(match[6])

			result := BenchmarkResult{
				Name:        name + match[2],
				Depth:       depth,
				Iterations:  iterations,
				TimePerOp:   timePerOp,
				BytesPerOp:  bytesPerOp,
				AllocsPerOp: allocsPerOp,
			}
			results = append(results, result)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	// Print parsed results for verification
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}

	// Get a names list without duplicates
	resultsSeparated := make(map[string][]BenchmarkResult)
	reNameWithoutDepth := regexp.MustCompile(`(\w+)-.*`)
	for _, result := range results {
		match := reNameWithoutDepth.FindStringSubmatch(result.Name)
		if match != nil {
			resultsSeparated[match[1]] = append(resultsSeparated[match[1]], result)
		}
	}

	// Create separate plots for each metric
	dirName := "plots"
	for r := range resultsSeparated {
		// Plot Time per Operation
		createPlot(resultsSeparated[r], "Time per Operation", "Time (ns/op)", dirName+"/"+r+"-time_per_op.png", func(result BenchmarkResult) float64 {
			return float64(result.TimePerOp)
		})

		// Plot Bytes per Operation
		createPlot(resultsSeparated[r], "Bytes per Operation", "Bytes (B/op)", dirName+"/"+r+"-bytes_per_op.png", func(result BenchmarkResult) float64 {
			return float64(result.BytesPerOp)
		})

		// Plot Allocations per Operation
		createPlot(resultsSeparated[r], "Allocations per Operation", "Allocations (allocs/op)", dirName+"/"+r+"-allocs_per_op.png", func(result BenchmarkResult) float64 {
			return float64(result.AllocsPerOp)
		})
	}
}

// createPlot generates a plot for a specific metric
func createPlot(results []BenchmarkResult, title, ylabel, filename string, valueFunc func(BenchmarkResult) float64) {
	values := make(plotter.Values, len(results))
	names := make([]string, len(results))

	for i, result := range results {
		values[i] = valueFunc(result)
		names[i] = result.Name
	}

	p := plot.New()

	p.Title.Text = title
	p.Y.Label.Text = ylabel

	w := vg.Points(20)

	bar, err := plotter.NewBarChart(values, w)
	if err != nil {
		panic(err)
	}

	p.Add(bar)
	p.NominalX(names...)

	if err := p.Save(10*vg.Inch, 6*vg.Inch, filename); err != nil {
		panic(err)
	}

	fmt.Printf("Plot saved to %s\n", filename)
}
