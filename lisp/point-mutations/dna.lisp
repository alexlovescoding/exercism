(defpackage #:dna
  (:use #:cl)
  (:export #:hamming-distance))
(in-package #:dna)


(defun hamming-distance (dna1 dna2)
  (cond ((string= dna1 dna2) 0)
    ((/= (length dna1) (length dna2)) nil)
    (t (setq count 0)
      (loop for i from 0 to (- (length dna1) 1)
        do (if (not (eq (char dna1 i) (char dna2 i)))
          (setq count (+ 1 count)))) count)))
