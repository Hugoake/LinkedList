// Hugo Åkesson, implements a linked list with data type interface{},
// which therefore supports any data type.

package genericLinkedList

import (
        "github.com/pkg/errors"
)

// A list element that stores a value of type T.
type listElement struct {
    data interface{}
    nextPtr *listElement
}


 // Returns the data in the listElement "depth"-steps forward.
//  If no such listElement exists, returns nil
func (e listElement) retrieve(depth int) interface{} {
    if depth < 0 {
        panic( errors.New("retrieveCalledWithNegativeDepth") )
    }
    
    if depth == 0 {
        return e.data
    } else if e.nextPtr == nil {
        return nil
    } else {
        return (*e.nextPtr).retrieve(depth-1)
    }
}


// A singly linked list of elements of type T.
type LinkedList struct{
    firstPtr    *listElement  // first element in list
    lastPtr     *listElement //  last element in list
    size        int         //   number of elements in list
}


// Create a new empty list.
func New() *LinkedList {
    return new(LinkedList)
}


// Insert the given element at the beginning of this list.
func (lst *LinkedList) AddFirst(element interface{}){
    lst.firstPtr = &listElement{element, lst.firstPtr}
    if lst.size + 1 == 1 {
        lst.lastPtr = lst.firstPtr // In a one element list, first = last
    }
    lst.size++
}


// Insert the given element at the end of this list.
func (lst *LinkedList) AddLast(element interface{}) {
    if lst.size == 0 {
        lst.AddFirst(element)
    } else { 
        (*lst.lastPtr).nextPtr = &listElement{element, nil} // Appending new element
        lst.lastPtr = (*lst.lastPtr).nextPtr // Updating which element is last
        lst.size++
    }
}


 // Return the first element of this list.
//  Return null if the list is empty.
func (lst LinkedList) GetFirst() interface{} {
    if lst.size == 0 {
        return nil 
    } else {
        return (*lst.firstPtr).data 
    }
}


 // Return the last element of this list.
//  Return null if the list is empty.
func (lst LinkedList) GetLast() interface{} {
    if lst.size == 0 {
        return nil 
    } else {
        return (*lst.lastPtr).data 
    }
}


 // Return the element at the specified position in this list.
//  Return null if index is out of bounds.
func (lst LinkedList) Get(index int) interface{} {
    if lst.size == 0 || index < 0 {
        return nil
    } else {
        return (*lst.firstPtr).retrieve(index)
    }
}


 // Remove and returns the first element from this list.
//  Return null if the list is empty.
func (lst *LinkedList) RemoveFirst() interface{} {
    
    if lst.size == 0 {
        return nil
    }
    old_data :=  (*lst.firstPtr).data
    lst.firstPtr = (*lst.firstPtr).nextPtr // Second element becomes first
    
    if lst.size == 1 {
        lst.lastPtr = nil   // If the list *had* length one, first = last
                           // => last should be set to nil as well
    }
    lst.size--
    
    return old_data
}


// Remove all elements from this list.
func (lst *LinkedList) Clear() {
    *lst = *New()
}


// Return the number of elements in this list.
func (lst LinkedList) Size() int {
    return lst.size
}


  //Return a string representation of this list.
 // The elements are enclosed in square brackets ("[]").
//  Adjacent elements are separated by ", ".
func (lst LinkedList) String() string {
    if lst.size == 0 {
        return "[]"
    } else {
        return "[" + (*lst.firstPtr).presentWithNext(", ") + "]"
    }
}


// Used by String()
func (e listElement) presentWithNext(delim string) string {
    if e.nextPtr == nil {
        return fmt.Sprintf("%v", e.data) // "Base case" - last element reached
    } else {
        return fmt.Sprintf("%v", e.data) + delim + (*e.nextPtr).presentWithNext(delim)
    }
}


 // Determines weather or not the LinkedList is in a correct "healthy" state
//  An improvement that should be made: Also return what kind of error
func (lst LinkedList) Healthy() bool {
    if lst.firstPtr != nil {
        // Follow and count the chain of listElements
        MAX_SIZE := 10000
        tally := 1
        for nextPtr := (*lst.firstPtr).nextPtr ;
            nextPtr != nil;
            nextPtr = (*nextPtr).nextPtr {
            
            tally++
            if tally == MAX_SIZE {
                fmt.Println("List to large to check (more elements than 10000)")
                return false
            }
        }
        if tally != lst.size {
            fmt.Println("Incorrect size")
            return false
        }
    }
    
        
    if lst.size == 0 {
        // If list is empty, neither first nor last should point anywhere
        if lst.firstPtr != nil || lst.lastPtr != nil {
            fmt.Println("Bad empty list")
            return false
        }
    } else {
        // Check that lastPtr and firstPtr points to listElement:s
        _, firstTypeGood := interface{}(*lst.firstPtr).(listElement)
        _, lastTypeGood := interface{}(*lst.lastPtr).(listElement)
        if !firstTypeGood || !lastTypeGood {
            fmt.Println("Bad first or last type")
            return false
        }
        
        // Check that last element doesn't point further
        if (*lst.lastPtr).nextPtr != nil {
            fmt.Println("Last pointer points")
            return false
        }
    }
    // Everything seems good
    return true
}