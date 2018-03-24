package uses

//indexAssigneNewPosition ...
func indexAssignNewPosition(workArray []int, oldIndex int, newIndex int) []int {
	//Given a slice and an index (old and new), allow a value define by arr[oldIndex] to be assigneda new position at newIndex
	//(0123, from [1] to [3] makes it 0231)
	workSize := len(workArray)
	if oldIndex > (workSize-1) || newIndex > (workSize-1) {
		return workArray //We can't proceed with this function
	}

	if oldIndex == newIndex {
		return workArray //Done before we even started
	}

	elementToMove := workArray[oldIndex]

	if newIndex > oldIndex { //Forward moving swap
		for i := oldIndex; i < newIndex; i++ {

			if workArray[i+1] == workArray[newIndex] {
				shiftDown := workArray[i+1] //One last shift
				workArray[i] = shiftDown

				workArray[i+1] = elementToMove //Finally putting the element in it's position
				break
			}

			shiftDown := workArray[i+1] //Shift down and replace, we have a copy sort of
			workArray[i] = shiftDown

		}
	} else if newIndex < oldIndex { //Backward moving swap
		for i := oldIndex; i > newIndex; i-- {

			if workArray[i-1] == workArray[newIndex] {
				shiftUp := workArray[i-1] //One last shift
				workArray[i] = shiftUp

				workArray[i-1] = elementToMove //Finally putting the element in it's position
				break
			}

			shiftUp := workArray[i-1] //Shift up and replace, we have a copy sort of
			workArray[i] = shiftUp

		}
	}

	return workArray

}
