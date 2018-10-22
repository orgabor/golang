/*

Slices, like arrays, also represent a list of elements.
The code to update and access a slice's elements looks much the same as the code to work with an array's elements.
 But slices a little easier to work with than arrays, once you know the proper techniques.
 So slices are more commonly used than arrays.

A slice doesn't actually store anything itself.
It references a portion of an underlying array that actually holds the data.
Most of the time, you won't have to worry about the underlying array.
But it can cause some suprising bugs if you don't know how it works.

We'll start by creating an array, and then basing some slices off of it.
This may not be the easiest way to use slices in your programs, but it does demonstrate what's going on.

Here we have an array with 5 elements.
We'll create 2 slices based on it.
The first will include index 0 of the array up until index 3 (but not including index 3).
The second will include index 2 up until index 5 (but not including index 5, which is a good thing, because the array doesn't have an index 5).

*/

package main

import "fmt"

func main() {
	a := [5]int{0, 1, 2, 3, 4}
	s1 := a[0:3]
	s2 := a[2:5]

	fmt.Println(a, s1, s2)
	changeArray()
	changeSliceElement()
	changeSliceSlice()
	sliceFunc()
	sliceAppend()
	sliceAppendV2()
	sliceSecurity()
	sliceForRange()
}

/*
f we run this, it will print the underlying array, as well as the two slices: [0 1 2 3 4] [0 1 2] [2 3 4]. Notice that the number at index 2 in the underlying array appears in both slices.

Now, let's try modifying that number in the underlying array. Then we'll print the array and both slices again.



*/

func changeArray() {
	a := [5]int{0, 1, 2, 3, 4}
	s1 := a[0:3]
	s2 := a[2:5]
	a[2] = 88
	fmt.Println(a, s1, s2)
}

/*
f we run it again, we'll see that the underlying array has been updated. But it also appears that the contents of both slices have changed. That's because, as we mentioned before, slices don't actually hold any data themselves. They're just a sort of window into the contents of the underlying array.

We can use slices to update the underlying array as well. Let's update the element at index 0 of the second slice, which is actually the element at index 2 of the underlying array. Then we'll print the array and both slices again.
*/
func changeSliceElement() {
	a := [5]int{0, 1, 2, 3, 4}
	s1 := a[0:3]
	s2 := a[2:5]
	s2[0] = 999
	fmt.Println(a, s1, s2)
}

/*
Run it again, and we'll see that the element at index 2 of the underlying array has been updated,
 and so both slices reflect the change as well.

We can re-slice a slice to reveal more of the underlying array.
*/

func changeSliceSlice() {
	a := [5]int{0, 1, 2, 3, 4}
	s1 := a[0:3]
	s2 := a[2:5]
	s1 = s1[0:4]
	fmt.Println(a, s1, s2)
}

/*
+++++++++++++++++++++++
If we try the same with the second slice, though, we'll get a runtime panic:
 s2 = s2[0:4]
  fmt.Println(a, s1, s2)
*/

/*
There are a couple functions useful for getting information about slices.
-len gets the current length of the slice, just as it does when used with an array.
-cap shows the capacity of the slice, which is usually the number of elements between
the start of the slice and the end of the underlying array.

Notice s1's capacity is higher than s2's.
s2 doesn't have any room to grow because it's right at the end of the underlying array.

*/

func sliceFunc() {
	a := [5]int{0, 1, 2, 3, 4}
	s1 := a[0:3]
	s2 := a[2:5]
	fmt.Println("len(s1):", len(s1), "cap(s1):", cap(s1))
	fmt.Println("len(s2):", len(s2), "cap(s2):", cap(s2))
}

/*
So what can we do if we need to add a value to a slice,
but we're at the end of the underlying array?
Go offers a built-in function called append that can add new values to a slice,
even if it's already at its capacity.
Let's call append with our s2 slice, and append the value 5 to it.
ow, append doesn't actually modify the existing slice,
 it returns a new slice with the same contents as the old one,
 plus the new values appended at the end.
 So we'll need to assign that return value back to the s2 variable.
 Then, we'll print everything out again so we can look at the changes.

*/

func sliceAppend() {
	a := [5]int{0, 1, 2, 3, 4}
	s1 := a[0:3]
	s2 := a[2:5]
	fmt.Println(a, s1, s2)
	fmt.Println("len(s1):", len(s1), "cap(s1):", cap(s1))
	fmt.Println("len(s2):", len(s2), "cap(s2):", cap(s2))
	s2 = append(s2, 5)
	fmt.Println(a, s1, s2)
	fmt.Println("len(s1):", len(s1), "cap(s1):", cap(s1))
	fmt.Println("len(s2):", len(s2), "cap(s2):", cap(s2))

}

func sliceAppendV2() {

	a := [5]int{0, 1, 2, 3, 4}
	s1 := a[0:3]
	s2 := a[2:5]
	s2[0] = 88
	s2 = append(s2, 5)
	fmt.Println(a, s1, s2) //extend s1 s2 AND the array
	s2[0] = 999
	fmt.Println(a, s1, s2) //extends s2 ONLY created a new underlining array
}

/*
If we run this, we'll see that s2's capacity is now
even higher than s1's because it's pointing at a bigger array.
And even though we appended just one element,
s2's capacity increased by 3, not by 1. That's because allocating an array is slow,
so append gives us more room than we've asked for, in case we're going to do more appends later.
If append runs out of capacity again, it will allocate an even bigger underlying array next time.

I mentioned before that creating an array and then basing slices off
of it may not be the easiest way to use slices in your programs.
It's also not the safest. Now that you have a better understanding of how slices work,
I'm going to show you a better way.

We showed you before how to pre-populate an array using curly braces: a := [3]int{1, 2, 3}.
You can create a pre-populated slice using almost the same notation;
you just leave off the size within the square brackets: s := []int{1, 2, 3}.
An underlying array will automatically be created for the slice,
so you don't have to worry about where to store it, or worry about accidentally altering it.
If you don't want to specify starting element values, you can leave those out: s := []int{}.

Once you have the slice, you can then append values as needed: s = append(s, 4, 5), s = append(s, 6, 7, 8),
and so on. We can print the slice's contents: fmt.Println(s), and then run it,
 and we'll see that all the appended values have accumulated in the slice.
 We don't have to worry about the underlying array at all. It just works.
 Until you have more specific needs,
 you'll probably find that this is the best way to create lists of items in Go.





*/

func sliceSecurity() {
	a := [3]int{1, 2, 3} //array
	s := []int{1, 2, 3}  //slice with underlining array
	s = append(s, 5, 6, 99, 7)
	s = append(s, 34, 56, 444, 98)
	fmt.Println(a, s)
}

/*ust as with arrays, you can use a for ... range loop to process each element in a slice:*/
func sliceForRange() {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		fmt.Println("value:", v)
	}

}
