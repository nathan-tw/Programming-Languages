-- 1. (a)
power1 :: Int -> Int -> Int
power1 n k = product [n | _ <-[1..k]]

-- 1. (b)
power2 :: Int -> Int -> Int
power2 _ 0 = 1
power2 n k  = let sqr i = i * i
                  half_k = k `div` 2  
                  sqrHalf = sqr ( power2 n half_k )
              in if even k then sqrHalf else n * sqrHalf

-- 2. (a)
myButLast :: [a] -> a
myButLast [x, _] = x
myButLast (x:xs) = myButLast xs

-- 2. (b)
rev2 :: [a] -> [a]
rev2 [x, y] = y:[x]
rev2 (x:xs) = x:xs

-- 3. (a)
tailUpto :: Int -> Int -> [Int] -> [Int]
tailUpto n k ls = if n>k then ls else tailUpto n (k-1) (k:ls)

-- 3. (b)
tailFib :: Int -> Int -> Int -> Int
tailFib = undefined

-- 4. (a)
palindrome :: [Int] -> Bool
palindrome [] = True
palindrome [_] = True
palindrome xs = if head xs == last xs then palindrome (init (tail xs)) else False

-- 4. (b)
removeOnce :: Int -> [Int] -> [Int]
removeOnce _ [] = [] 
removeOnce a (b:bc) | a == b    = bc 
                    | otherwise = b : removeOnce a bc
isPermutation :: [Int] -> [Int] -> Bool
isPermutation [] [] = True
isPermutation xs ys| length xs /= length ys = False
isPermutation (x:xx) ys = isPermutation xx (removeOnce x ys)