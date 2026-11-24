# Single Number

> NeetCode 150 · Bit Manipulation — XOR all values: x^x = 0 and order never matters, so every pair cancels and the loner remains. Go, stdlib only, tested.

## Problem

Every value appears exactly twice except one; find it in O(n) time and O(1) space — the constraint that rules out sets and sorting.

## Approach

Three algebraic facts make a one-liner: XOR is commutative and associative (pairs can be mentally rearranged to sit together), x^x = 0 (each pair annihilates), and x^0 = x (the loner passes through). Folding the whole array through XOR therefore leaves exactly the single value — sign and zero included, no cases. This cancellation trick is the seed of a family: repeat-and-missing (dsa-011) splits XOR buckets by a distinguishing bit, and "single among triples" upgrades the idea to per-bit counting mod 3.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
