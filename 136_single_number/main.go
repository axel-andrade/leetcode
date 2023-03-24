package main

import "fmt"

/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.



Example 1:

Input: nums = [2,2,1]
Output: 1
Example 2:

Input: nums = [4,1,2,1,2]
Output: 4
Example 3:

Input: nums = [1]
Output: 1


Constraints:

1 <= nums.length <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
Each element in the array appears twice except for one element which appears only once.
*/

/*
Abordagens:
Brute Force
Intuition:
Itere através de cada elemento nos nums e verifique se algum elemento não aparece duas vezes, nesse caso retorne o elemento.
Tempo: O(n^2)
Espaço: O(1)

Use
a Intuição de Classificação:
Se os elementos do array nums estiverem ordenados/quando os classificarmos, podemos comparar os vizinhos para encontrar o único elemento. Já foi mencionado que todos os outros elementos aparecem duas vezes, exceto um.
Tempo: O(nlogn) para ordenar e depois O(n) para verificar os elementos vizinhos
Espaço: O(1)

Use Hashing/Set
Intuition:
i) À medida que iteramos pelo array nums, armazenamos os elementos encontrados e verificamos se os encontramos novamente enquanto a iteração continua. Ao verificar se os encontramos novamente, mantemos um objeto/variável single_element que armazena esse único elemento, eventualmente retornando o single_element.
ii) A outra maneira é manter um hashmap/dicionário num_frequency e iterar sobre ele para encontrar qual tem exatamente 1 frequência e retornar essa chave/num.
Tempo: O(n) para iteração no array nums
Espaço: O(n) para hash

	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	var num int

	for k, v := range m {
		if v == 1 {
			num = k
			break
		}
	}

	return num

Use Xor/ Intuição de Manipulação de Bits
: Xor de quaisquer dois num dá a diferença de bit como 1 e mesmo bit como 0.
Assim, usando isso, obtemos 1 ^ 1 == 0 porque os mesmos números têm os mesmos bits.
Portanto, sempre obteremos o elemento único porque todos os mesmos serão avaliados como 0 e 0^single_number = single_number.
Tempo: O(n)
Espaço: O(1)


*/
func singleNumber(nums []int) int {
	var xor int = 0

	for _, v := range nums {
		xor ^= v
	}

	return xor
}

func main() {
	inputs := [][]int{
		{2, 2, 1},
	}

	for _, v := range inputs {
		fmt.Println(singleNumber(v))
	}

}
