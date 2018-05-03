def arraySwap(arr, flag):
    array = arr
    length = len(array)
    for i in range(array):
        if array[i] == flag and array[length-i - 1] != flag:
            array[i], array[length-i-1] = array[length-i-1], array[i]
        elif array[i] == flag and array[length-i-1] == flag:
            length -= 1
            i -= 1
    
    return array