# validator

A simple form validator made with Go.

## Installation

```Shell
go get -u github.com/ezlosswm/validator
```

## Usage

```Go
firstname := "ezlos"
if ok := IsEmpty(firstname); !ok {
	log.Printf("(%q) value provided is empty.\n", firstname)
	return
}

a, b := "hello", "world"
if ok := IsMatching(a, b); !ok {
	log.Printf("%s is not equal to %s.\n", a, b)
    return
}

email := "ezlos@example.com"
if ok := CheckEmail(email); !ok {
	log.Printf("%s not a valid email.\n", email)
	return
}

password := "Password123"
if err := CheckLength(password); err != nil {
	log.Println(err)
	return
}
```

## Todo 
- [ ] Test for password strenght.
