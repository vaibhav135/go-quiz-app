package utility 


func Compare(element string, slice []string) bool {
  for _, i := range slice {
    if element == i {
      return true
    }
  } 

  return false
}
