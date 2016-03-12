module Grains(square, total) where

square :: Integer -> Integer
total :: Integer

square a = 2^(a-1)
total = 2^64-1
