module DNA (toRNA) where

import Data.Maybe

rnaMap :: [(Char, Char)]
rnaMap = [('G', 'C'), ('C', 'G'), ('T', 'A'), ('A', 'U')]

toRNA :: String->String
toRNA dna = [Data.Maybe.fromJust $ lookup s rnaMap | s <- dna]
