module Gigasecond(fromDay) where
import Data.Time.Clock(addUTCTime, UTCTime, NominalDiffTime)

fromDay :: UTCTime -> UTCTime
fromDay t = addUTCTime (10^9 :: NominalDiffTime) t
