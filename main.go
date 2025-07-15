package main

import "fmt"

// domain parameter variables. these are chosen and can be public / known to the users that will be signing and verifying signature
//	q = the field size
//	FR = an indication of the basis used
//	h = the cofactor
//	n = order of the point G, we will be choosing a number with 512 bits since it has the strength of a 256-bit num
// 		this also means we have to use SHA-512, recommendation from the paper
//	Type = elliptical curve model used
//	a = param part of the curve equation
//	b = param part of the curve equation
//	G = base point of the prime order on the curve: G = (xG, yG)
//	domain_parameter_seed = random bit string (optional)

// other notes i guess (i need to brush up on this) Galois Field
//ECDSA and deterministic ECDSA are defined for two arithmetic fields: the finite field GF(p)
//and the finite field GF(2𝑚𝑚). For the field GF(p), p is required to be an odd prime.

func main() {
	fmt.Println("starting dev setup")
}
