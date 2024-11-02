package vectozavr

import "math"

type Matrix struct {
	m [4][4]float64
}

//Создание новой матрицы
func NewMatrix(m [4][4]float64) Matrix {
	return Matrix{
		m: m,
	}

}

//Умножение матрицы на матрицу
func (m Matrix) MatMul(n Matrix) Matrix {
	var result [4][4]float64
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				result[i][j] += m.m[i][k] * n.m[k][j]
			}

		}

	}
	return NewMatrix(result)
}

//Умножение матрицы на вектор размером 4
func (m Matrix) Vec4Mul(v Vec4) Vec4 {
	var v1 Vec4
	v1.X = m.m[0][0]*v.X + m.m[0][1]*v.Y + m.m[0][2]*v.Z + m.m[0][3]*v.W
	v1.Y = m.m[1][0]*v.X + m.m[1][1]*v.Y + m.m[1][2]*v.Z + m.m[1][3]*v.W
	v1.Z = m.m[2][0]*v.X + m.m[2][1]*v.Y + m.m[2][2]*v.Z + m.m[2][3]*v.W
	v1.W = m.m[3][0]*v.X + m.m[3][1]*v.Y + m.m[3][2]*v.Z + m.m[3][3]*v.W

	return v1

}

//Умножение матрицы на вектор размером 3
func (m Matrix) Vec3Mul(v Vec3) Vec3 {
	var v1 Vec3
	v1.X = m.m[0][0]*v.X + m.m[0][1]*v.Y + m.m[0][2]*v.Z
	v1.Y = m.m[1][0]*v.X + m.m[1][1]*v.Y + m.m[1][2]*v.Z
	v1.Z = m.m[2][0]*v.X + m.m[2][1]*v.Y + m.m[2][2]*v.Z

	return v1
}

//Единичная матрица
func Identity() Matrix {
	return NewMatrix([4][4]float64{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
}

//Матрица, все элементы который заданные числа
func Constant(value float64) Matrix {
	var m Matrix
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m.m[i][j] = value
		}
	}
	return m
}

//Пустая матрица (все значения нулевые)
func ZeroMatrix() Matrix {
	return NewMatrix([4][4]float64{})
}

//Матрица изменения масштаба
func Scale(v Vec3) Matrix {
	return NewMatrix([4][4]float64{
		{v.X, 0, 0, 0},
		{0, v.Y, 0, 0},
		{0, 0, v.Z, 0},
		{0, 0, 0, 1},
	})

}

//Матрица перемещения
func Translation(v Vec3) Matrix {
	return NewMatrix([4][4]float64{
		{1, 0, 0, v.X},
		{0, 1, 0, v.Y},
		{0, 0, 1, v.Z},
		{0, 0, 0, 1},
	})

}

//Матрица поворота вокруг оси X
func RotationX(angle float64) Matrix {
	c := math.Cos(angle)
	s := math.Sin(angle)

	return NewMatrix([4][4]float64{
		{1, 0, 0, 0},
		{0, c, -s, 0},
		{0, s, c, 0},
		{0, 0, 0, 1},
	})

}

//Матрица поворота вокруг оси Y
func RotationY(angle float64) Matrix {
	c := math.Cos(angle)
	s := math.Sin(angle)

	return NewMatrix([4][4]float64{
		{c, 0, s, 0},
		{0, 1, 0, 0},
		{-s, 0, c, 0},
		{0, 0, 0, 1},
	})

}

//Матрица поворота вокруг оси Z
func RotationZ(angle float64) Matrix {
	c := math.Cos(angle)
	s := math.Sin(angle)

	return NewMatrix([4][4]float64{
		{c, -s, 0, 0},
		{s, c, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
}

//Матрица поворота (по всем осям)
func (m Matrix) Rotation(v Vec3) Matrix {
	r := RotationX(v.X).MatMul(RotationY(v.Y))
	return r.MatMul(RotationZ(v.Z))
}
