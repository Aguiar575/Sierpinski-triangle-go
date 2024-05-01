# Sierpinski Triangle Generator

This program generates a Sierpinski triangle, a fractal named after the Polish mathematician Wacław Sierpiński. The Sierpinski triangle is a fascinating mathematical object with intricate self-similar patterns.

## Sierpinski Triangle

The Sierpinski triangle is a fractal and a classic example of self-similar patterns in mathematics. It is formed by recursively subdividing an equilateral triangle into smaller triangles and removing the central triangle at each step. The resulting shape is a complex and beautiful pattern that exhibits self-similarity on different scales.

## How it Works

1. **Initialization**: The program starts by defining a base unit (usually represented by an asterisk `*`) surrounded by spaces. This forms the initial triangle.

2. **Iteration**: It then iteratively builds upon this initial triangle. In each iteration, it adds smaller triangles (each consisting of the base unit) to the top of the existing triangle. These smaller triangles are centered above the previous triangle.

3. **Recursion**: The process repeats recursively, with each iteration adding smaller triangles on top of the previous ones. This recursive pattern creates the intricate structure of the Sierpinski triangle.

4. **Visualization**: After each iteration, the program prints the current state of the triangle. By printing the triangles one after another with slight delays, it creates an animation effect, allowing the viewer to observe the fractal's formation.

## Running the Program

To run the program, simply execute the main Go file. 
```bash
go run main.go
```




