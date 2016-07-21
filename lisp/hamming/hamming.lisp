(defpackage #:hamming
  (:use #:cl)
  (:export #:distance))
(in-package #:hamming)


(defun distance (dna1 dna2)
  "Compute the Hamming distance between two DNA strands"
  (when (= (length dna1) (length dna2))
    (loop for char1 across dna1
      for char2 across dna2
      count(not (char= char1 char2)))))