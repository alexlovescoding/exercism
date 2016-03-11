module Sublist(Sublist(..), sublist) where

import Data.List (isInfixOf)

data Sublist = Equal | Sublist | Superlist | Unequal
  deriving(Eq, Show)

sublist :: Eq a => [a]->[a]->Sublist
sublist as bs | as == bs          = Equal
              | as `isInfixOf` bs = Sublist
              | bs `isInfixOf` as = Superlist
              | otherwise         = Unequal
