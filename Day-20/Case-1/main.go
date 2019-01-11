package main

import(
	"os"
	"bufio"
	"regexp"
	"strconv"
	"fmt"
	"math"
)

type particle struct {
	number int
	posX int
	posY int
	posZ int
	accX int
	accY int
	accZ int
	velX int
	velY int
	velZ int
}

func main() {

	file, _ := os.Open("input.txt")
	input := bufio.NewScanner(file)
	particles := []particle{}
	particleNumber := 0

	for input.Scan() {

		particleRegex := regexp.MustCompile("p=<(-?\\d*),(-?\\d*),(-?\\d*)>, v=<(-?\\d*),(-?\\d*),(-?\\d*)>, a=<(-?\\d*),(-?\\d*),(-?\\d*)>")
		particleValues := particleRegex.FindStringSubmatch(input.Text())
	
		newParticle := particle {
			number: particleNumber,
			posX: convertToInt(particleValues[1]),
			posY: convertToInt(particleValues[2]),
			posZ: convertToInt(particleValues[3]),
			accX: convertToInt(particleValues[4]),
			accY: convertToInt(particleValues[5]),
			accZ: convertToInt(particleValues[6]),
			velX: convertToInt(particleValues[7]),
			velY: convertToInt(particleValues[8]),
			velZ: convertToInt(particleValues[9]),
		}
		particles = append(particles, newParticle)

		particleNumber++

	}

	minTotalVel := float64(2147483647)
	var closestParticle particle

	for _, particle := range particles {

		totalVel := math.Abs(float64(particle.velX)) + math.Abs(float64(particle.velY)) + math.Abs(float64(particle.velZ))
		if totalVel < minTotalVel {
			minTotalVel = totalVel
			closestParticle = particle
		}

	}

	fmt.Println(closestParticle.number)

}

func convertToInt(s string) (result int) {
	result, _ = strconv.Atoi(s)
	return result
}