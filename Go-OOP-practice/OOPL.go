package main

import "fmt"
import "strconv"


// Q1
type Point struct{
  x int
  y int
}
var count int = 0
// constructor
func newPoint(x int, y int) Point {
    p := Point{x,y}
    count+=1
    return p
}

// instance methods
func (p Point) getX()  int{
  return p.x
}

func (p Point) getY() int{
  return p.y
}

func (p *Point) move(dx int, dy int){
  p.x += dx
  p.y += dy
}

func (p Point) toString() string{
  return "A point at "+strconv.Itoa(p.getX())+","+strconv.Itoa(p.getY())
}

func (p Point) draw() {
  fmt.Println(p.toString())
}

//------------------------------------------------------

//Q2
type ColorPoint struct{
  Point
  color_ string `default:"WHITE"`
}

//colorpoint's constructor
func newColorPoint(x int, y int, color string) ColorPoint {
    c := ColorPoint{Point{x,y}, color}
    return c
}

func (c ColorPoint) getColor() string{
  return c.color_
}

func (c ColorPoint) setColor(color string){
  c.color_ = color
}

// method overide
func (c ColorPoint) toString()string{
  return c.Point.toString()+" with color "+c.getColor()
}

func (c ColorPoint) draw() {
  fmt.Println(c.toString())
}


func main() {
  fmt.Println("Q1\n")
  p1 := newPoint(3,5)
  p1.draw()

  fmt.Println(count)

  p1.move(1, 1)
  p1.draw()

  p2 := newPoint(2,2)
  p2.draw()

  fmt.Println(count)
  fmt.Println()

  fmt.Println("Q2")
  cpt := newColorPoint(3, 5, "RED")
  cpt.draw()
  cpt.move(1, 1)
  cpt.draw()


}

//------------------------------------------------------

//Q3

