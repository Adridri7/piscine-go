package piscine 

func Rot14(s string) (res string) {
  for _, char := range s{
    if char == ' ' || char == '!' || char == '?' || char >= '0' && char <= '9'{
      res += string(char)
    } else{
      if char >= 'a' && char <= 'z'{
        if char + 14 > 'z'{
          diff :=  char - 26 + 14
          res += string(diff)
        } else{
          res += string(char + 14)
        }
      } else if char >= 'A' && char <= 'Z' {
        if char + 14 > 'Z'{
          diff := char + 14 - 26
          res += string(diff)
        } else{
          res += string(char + 14)
        }
      }
    }
  }
  return res
}