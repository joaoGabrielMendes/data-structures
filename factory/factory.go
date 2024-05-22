package factory

type Array interface {
        Push(element interface{}) int                                                                           // Add element on top ✅
        Pop() (interface{}, error)                                                                              // Remove the last element ✅
        Shift() (interface{}, error)                                                                            // Remove the first element ✅
        Unshift(elements ...interface{}) int                                                                    // Add elements at the first place ✅
        Concat(arrays ...Array) Array                                                                           // Concatenate arrays
        Slice(start, end int) Array                                                                             // Return a shallow copy of a portion of an array
        Splice(start, deleteCount int, items ...interface{}) Array                                              // Add/remove elements from the array
        ForEach(callback func(element interface{}, index int))                                                  // Execute a function for each element
        Map(callback func(element interface{}, index int) interface{}) Array                                    // Create a new array with the results of calling a function for each element
        Filter(callback func(element interface{}, index int) bool) Array                                        // Create a new array with all elements that pass the test
        Reduce(callback func(acc, element interface{}, index int) interface{}, initial interface{}) interface{} // Apply a function to reduce the array to a single value
        Find(callback func(element interface{}, index int) bool) (interface{}, error)                           // Return the first element that satisfies the test
        FindIndex(callback func(element interface{}, index int) bool) int                                       // Return the index of the first element that satisfies the test
        IndexOf(element interface{}) int                                                                        // Get the first index of the element
        LastIndexOf(element interface{}) int                                                                    // Get the last index of the element
        Includes(element interface{}) bool                                                                      // Check if the array includes the element
        Some(callback func(element interface{}, index int) bool) bool                                           // Check if at least one element satisfies the condition
        Every(callback func(element interface{}, index int) bool) bool                                          // Check if all elements satisfy the condition
        Sort(less func(i, j interface{}) bool)                                                                  // Sort the array
        Reverse()                                                                                               // Reverse the array
        Size() int                                                                                              // Get the size of the array
        IsEmpty() bool                                                                                          // Check if the array is empty
}
