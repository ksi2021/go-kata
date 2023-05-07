# Easy total(26)


## Easy №1 N-th Tribonacci Number

```
func tribonacci(n int) int {
	if n == 4 {
		return 4
	}
	if n == 3 {
		return 2
	}
	if n == 2 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}

	return tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)
}

```

## Easy №2 Concatenation of Array

```
func getConcatenation(nums []int) []int {
	size := len(nums)
	ans := make([]int, size*2)

	copy(ans, nums)

	for i := size; i < size*2; i++ {
		ans[i] = nums[i-size]
	}

	return ans
}

```

## Easy №3 Convert the Temperature

```
func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.80 + 32.00}
}

```

## Easy №4 Build Array from Permutation

```
func buildArray(nums []int) []int {
	ans := make([]int, len(nums))

	for i := range nums {
		ans[i] = nums[nums[i]]
	}
	return ans
}

```

## Easy №5 Count of Matches in Tournament

```
func numberOfMatches(n int) int {
	matches := 0

	for n > 1 {
		if n%2 == 0 {
            n /= 2
			matches += n
			
		} else {
			n = n - 1
            n = n / 2
			matches += n
			n += 1
		}
	}

	return matches
}

```



## Easy №7 Defanging an IP Address

```
func defangIPaddr(address string) string {
	result := ""

	for _, i := range address {
		if string(i) == "." {
			result += "[.]"
			continue
		}

		result += string(i)
	}

	return result
}

```



## Easy №8 Kth Missing Positive Number

```
func findKthPositive(arr []int, k int) int {
	start := 1
	currentMiss := 0
	counter := 0

	for start != arr[0] {
		currentMiss = start
		start++
		counter++

		if counter == k {
			break
		}
	}

	for i := 1; i < len(arr); i++ {
		if counter == k {
			break
		}
		if arr[i]-1 != arr[i-1] {
			current := arr[i]
			prev := arr[i-1]
			for current-1 > prev {
				prev++
				currentMiss = prev
				counter++

				if counter == k {
					break
				}
			}
		}

		fmt.Println(counter, currentMiss, arr[i])

	}

	if counter < k {
		start := arr[len(arr)-1] + 1
		for counter != k {
			currentMiss = start
			start++
			counter++
		}
	}

	return currentMiss
}


```
## Easy №9 Final Value of Variable After Performing Operations

```
func finalValueAfterOperations(operations []string) int {
    x := 0
	
	for _,v := range operations{
		if v == "--X" || v == "X--"{
			x--
		}else{
			x++
		}
	}
	return x
}

```


## Easy №10 Shuffle the Array

```
func shuffle(nums []int, n int) []int {
	ans := make([]int, len(nums))
	counter := 0
	for i := 0; counter < n; i += 2 {

		ans[i] = nums[counter]
		ans[i+1] = nums[n+counter]
		counter++
	}

	return ans
}


```

## Easy №11 Running Sum of 1d Array

```
func runningSum(nums []int) []int {
	ans := make([]int, len(nums))

	ans[0] = nums[0]

	for i:= 1; i < len(nums); i++{
		ans[i] = nums[i] + ans[i - 1]
	}
	return ans
}


```

## Easy №12 Number of Good Pairs

```
func numIdenticalPairs(nums []int) (ans int) {
	temp := [101]int{}
	for _, x := range nums {
		ans += temp[x]
		temp[x]++
	}
	return
}

```

## Easy №13 Jewels and Stones

```
func numJewelsInStones(jewels string, stones string) int {
    counter := 0

    for _,v := range jewels{
        e := strings.Count(stones, string(v))
        if e != 0{
            counter+=e
        }
    }

    return counter
}

```

## Easy №14 Richest Customer Wealth

```
func maximumWealth(accounts [][]int) int {
	max :=0
	temp:=0
	for _,i :=range accounts{
		
		temp = 0
		for _,j := range i{
			temp+=j
		}

		if temp > max{
			max = temp
		}
	}

	return max
}

```

## Easy №15 Design Parking System

```
type ParkingSystem struct {
	B_pace_count  int
	M_place_count int
	S_place_count int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.B_pace_count > 0 {
			this.B_pace_count--
			return true
		}
	case 2:
		if this.M_place_count > 0 {
			this.M_place_count--
			return true
		}
	case 3:
		if this.S_place_count > 0 {
			this.S_place_count--
			return true
		}
	}

	return false
}

```


## Easy №16 Smallest Even Multiple

```
func smallestEvenMultiple(n int) int {
	for i := n; i <= 300; i++ {
		if i%2 == 0 && i%n == 0 {
			return i
		}
	}

	return 0
}

```

## Easy №17 Maximum Number of Words Found in Sentences

```
func mostWordsFound(sentences []string) int {
	max := 0
	temp := 0
	for _, v := range sentences {
		temp = len(strings.Split(v," "))
		if  temp > max{
			max = temp
		}
	}
	return max
}


```

## Easy №20 Kids With the Greatest Number of Candies

```
func kidsWithCandies(candies []int, extraCandies int) []bool {
    res := make([]bool,len(candies))

	max := 0 

	for _, v := range candies{
		if v > max{
			max = v
		}
	}

	for i := range candies{
		res[i] = candies[i] + extraCandies >= max 
	}

	return res
}


```

## Easy №21 Kids With the Greatest Number of Candies

```
func subtractProductAndSum(n int) int {
	sum := 0
	product := 1

	num := 0
	s_arr := strings.Split(strconv.Itoa(n), "")

	for _, v := range s_arr {
		num, _ = strconv.Atoi(v)
		sum += num
		product *= num
	}
	return product - sum

}

```

## Easy №22 How Many Numbers Are Smaller Than the Current Number

```
func smallerNumbersThanCurrent(nums []int) []int {
	size := len(nums)
	ans := make([]int , size)

	t := 0
	for i := 0 ; i < size ; i++{
		for j :=0 ; j < size; j++{
			if nums[i] > nums[j]{
				t++
			}
		}
		ans[i] = t
		t = 0
	}
	return ans

}

```


## Easy №23 Goal Parser Interpretation

```
func interpret(command string) string {
	words := map[string]string{"()": "o", "(al)": "al", "G": "G"}
	ans := ""

	for {
		for k, v := range words {
			index := strings.Index(command, k)
			if index == 0 {
				ans += v

				if len(k) == len(command) {
					return ans
				}
				command = command[len(k):]

			}
		}
	}
	return ans
}

```

## Easy №24 Decode XORed Array

```
func decode(encoded []int, first int) []int {
	ans := []int{first}
	for i, e := range encoded {
		ans = append(ans, ans[i]^e)
	}
	return ans
}

```


## Easy №25 Create Target Array in the Given Order

```
func createTargetArray(nums []int, index []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		copy(ans[index[i]+1:], ans[index[i]:])
		ans[index[i]] = v
	}
	return ans
}

```

## Easy №26 Decompress Run-Length Encoded List

```
func decompressRLElist(nums []int) []int {
	ans := []int{}

	for i := 1; i <= len(nums)-1; i += 2 {
		for j := 0; j < nums[i-1]; j++ {
			ans = append(ans, nums[i])
		}
	}

	return ans

}

```

## Easy №27 Split a String in Balanced Strings

```
func balancedStringSplit(s string) int {
	total := 0
	prev_index := 0
	size := len(s)
	for i := 0; i < size; i++ {
		if strings.Count(s[prev_index:i ], "L") == strings.Count(s[prev_index:i ], "R") {
			prev_index = i
			total++

			fmt.Println(s[prev_index:i], "L", "R", "i = ", i)
		}
	}
	return total
}

```

## Easy №28 Count the Digits That Divide a Number

```
func countDigits(num int) int {
	total := 0

	arr := strings.Split(strconv.Itoa(num), "")

	for _, v := range arr {
		val, _ := strconv.Atoi(v)
		if num%val == 0 {
			total++
		}
	}

	return total
}

```

## Easy №29 XOR Operation in an Array

```
func xorOperation(n int, start int) int {
	xor := start
	for i := 1; i < n; i++ {
		xor^= start + (2 * i)
	}
	return xor
}

```

# Medium total(16)


## Medium №12 Merge In Between Linked Lists

```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var left, right *ListNode
	counter := 0
	temp := list1
	for counter < a - 1 {
		temp = temp.Next
		counter++
	}
	left = temp

	for counter < b {
		temp = temp.Next
		counter++
	}
	right = temp.Next

	left.Next = list2

	for list2.Next != nil {
		list2 = list2.Next
	}

	list2.Next = right

	return list1
}

```

## Medium №14 XOR Queries of a Subarray

```
func xorQueries(arr []int, queries [][]int) []int {
   res := make([]int , len(queries))
   
   
   for k,v := range queries{
		for i:= v[0]; i <= v[1]; i++{
			res[k] ^= arr[i]
		}
   }
   return res
}

```

## Medium №15 Partitioning Into Minimum Number Of Deci-Binary Numbers

```
func minPartitions(n string) int {
	var max rune
	for _, ch := range n {
		if ch > max {
			max = ch
		}
	}
	return int(max - '0')
}

```

## Medium №16 Subrectangle Queries

```
type SubrectangleQueries struct {
    items [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
    return SubrectangleQueries{items : rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
    for i:= row1; i <= row2; i++{
        for j:= col1 ; j <= col2; j++{
            this.items[i][j] = newValue
        } 
    }
}


func (this *SubrectangleQueries) GetValue(row int, col int) int {
    return this.items[row][col]
}

```

## Medium №17 Queries on Number of Points Inside a Circle

```
func countPoints(points [][]int, queries [][]int) []int {
	ans := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		for j := 0; j < len(points); j++ {
			part1 := (points[j][0] - queries[i][0])
			part2 := (points[j][1] - queries[i][1])
			if (part1*part1)+(part2*part2) <= queries[i][2]*queries[i][2] {
				ans[i]++
			}
		}
	}

	return ans
}

```

## Medium №18 Group the People Given the Group Size They Belong To

```
func groupThePeople(groupSizes []int) [][]int {
    groups := make(map[int][]int)
    res := [][]int{}
    for i := 0; i < len(groupSizes); i++ {
        key := groupSizes[i]
        groups[key] = append(groups[key], i)

        if key == len(groups[key]){
            res = append(res,groups[key])
            groups[key] = []int{}
        }
    }
    
    return res
}

```

## Medium №20 Queries on a Permutation With Key (runtime beats 100%  [0ms])

```
func processQueries(queries []int, m int) []int {
	P := make([]int, m)
	ans := make([]int, len(queries))

	for i := 1; i <= m; i++ {
		P[i-1] = i
	}

	for k, v := range queries {
		idx := findIndex(P, v)
		ans[k] = idx

		left := P[:idx]
		right := P[idx+1:]

		P = []int{v}
		P = append(P, left...)
		P = append(P, right...)
	}

	return ans

}

func findIndex(arr []int, value int) int {

	for k, v := range arr {
		if v == value {
			return k
		}
	}
	return 0
}

```

## Medium №2 Sort the Students by Their Kth Score

```
func sortTheStudents(score [][]int, k int) [][]int {
    sort.Slice(score, func(i, j int) bool {
        return score[i][k] > score[j][k]
    })
    return score
}

```


## Medium №3 Merge Nodes in Between Zeros

```
func mergeNodes(head *ListNode) *ListNode {

	var start *ListNode

	start = head

	for start.Next != nil {

		if start.Val != 0 && start.Next.Val != 0 {
			start.Next.Val += start.Val
			start.Val = 0
		}
		start = start.Next
	}

	temp := head
	prev := head
	for temp.Next != nil {
		if temp.Val != 0 {
			prev.Next = temp
			prev = prev.Next
		} 
		temp = temp.Next
	}

	prev.Next = nil

	return head.Next

}

```

## Medium №5 Maximum Twin Sum of a Linked List

```
func pairSum(head *ListNode) int {
	var arr []int

	max := 0

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	n := len(arr)
	for i := 0; i < n/2; i++ {
		sum := arr[i] + arr[n-i-1]
		if sum > max {
			max = sum
		}
	}
	return max

}

```

## Medium №11 Delete Leaves With a Given Value [Runtime 0ms Beats 100%]

```
func removeLeafNodes(root *TreeNode, target int) *TreeNode {

	removeLeaf(root, target)
	if root != nil && root.Val == target && root.Left == nil && root.Right == nil{
		return nil
	}
	return root
}

func removeLeaf(node *TreeNode, val int) {
	if node == nil {
		return
	}

	removeLeaf(node.Left, val)
	removeLeaf(node.Right, val)

	if node.Left != nil && node.Left.Val == val && node.Left.Left == nil && node.Left.Right == nil {
		node.Left = nil
	}

	if node.Right != nil && node.Right.Val == val && node.Right.Right == nil && node.Right.Left == nil {
		node.Right = nil
	}
}

```


## Medium №13 Maximum Sum of an Hourglass

```
func maxSum(grid [][]int) int {
    var max int = -666

	for i:=0; i < len(grid) - 2; i++{
		for j :=0; j < len(grid[0]) - 2; j++{
			sum:= grid[i][j] + grid[i][j+1] + grid[i][j+2] + grid[i+1][j+1] + grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			
			if sum > max{
				max = sum
			}
		}
	}
	return max
}

```

## Medium №9 Arithmetic Subarrays

```
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	ans := make([]bool, len(l))
	var subArray []int
	for i := 0; i < len(ans); i++ {

		subArray = []int{}
		subArray = append(subArray, nums[l[i]:r[i] + 1]...)
		sort.Ints(subArray)
		ans[i] = true
		for j:=0 ; j < len(subArray) - 1; j++{
			if subArray[j+1] - subArray[j] != subArray[1] - subArray[0]{
				ans[i] = false
				break
			}
		}
	}

	return ans
}

```

## Medium №1 Deepest Leaves Sum

```
func deepestLeavesSum(root *TreeNode) int {
	temp := make(map[int][]int)
	sum := 0
	maxNumber := 0
	DeepLeaf(root, 0, temp, &maxNumber)

	for _, v := range temp[maxNumber] {
		sum += v
	}

  fmt.Println(temp)
	return sum
}

func DeepLeaf(node *TreeNode, level int, MAP map[int][]int, outLevel *int) {
	if node == nil {
		return
	}

	DeepLeaf(node.Left, level+1, MAP, outLevel)
	DeepLeaf(node.Right, level+1, MAP, outLevel)

	if node.Left == node.Right && level >= *outLevel {
		MAP[level] = append(MAP[level], node.Val)
		*outLevel = level
	}
}

```

## Medium №8 Minimum Number of Vertices to Reach All Nodes

```
func findSmallestSetOfVertices(n int, edges [][]int) []int {
    temp  := make(map[int]int)
    for _, edge := range edges {
        temp[edge[1]]++
    }
    result := make([]int,0,n)
    for i := 0; i < n; i++ {
        if _, ok := temp[i]; !ok {
            result = append(result, i)
        }
    }

    return result
}

```

## Medium №6 Maximum Binary Tree

```
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max, idx := maxVal(nums)
	root := &TreeNode{Val :max, Left: constructMaximumBinaryTree(nums[:idx]),Right: constructMaximumBinaryTree(nums[idx+1:])}
	return root
}

func maxVal(nums []int) (int, int) {

	max, idx := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			idx = i
		}
	}

	return max, idx
}

```



