### Function Types

    1. Anonymous Function
    2. Function Expression or Assign function in variable
    3. Higher Order function or first class function
    4. Callback function
    5. Variadic function
    7. Init function - you not call this, computer calls this
    8. Closure - close over
    9. Defer function - last in first out
    10. Receiver function
    11. IIFE = immediately invoked function expression

### 🔹 Struct কী?

Struct (structure) হলো Go-এর একটা composite data type,
যা একাধিক field (ভ্যারিয়েবল) একসাথে ধরে রাখতে পারে।

👉 সহজভাবে বললে:

Struct = একাধিক data কে একসাথে রাখার container।

এটা অনেকটা “object” এর মতো (যেমন JavaScript বা Python এ object/class থাকে)।

🔹 Basic Example
package main
import "fmt"

type Person struct {
name string
age int
}

func main() {
p1 := Person{name: "Rahim", age: 25}
fmt.Println(p1)
fmt.Println(p1.name)
fmt.Println(p1.age)
}

🔍 ব্যাখ্যা

1️⃣ Struct তৈরি করা হলো

type Person struct {
name string
age int
}

type কীওয়ার্ড দিয়ে নতুন একটা টাইপ তৈরি করা হয়।

Struct-এর মধ্যে field define করা হয় (name, age)।

2️⃣ Struct ব্যবহার করা হলো

p1 := Person{name: "Rahim", age: 25}

এখানে p1 হলো Person টাইপের একটা ভ্যারিয়েবল।

3️⃣ Field Access করা

fmt.Println(p1.name)
fmt.Println(p1.age)

Output:

{Rahim 25}
Rahim
25

🔹 Field গুলোর মান পরিবর্তন করা যায়
p1.age = 30
fmt.Println(p1)

Output:

{Rahim 30}

🔹 Anonymous Struct

যদি একবারের জন্য struct দরকার হয়, আলাদা type না বানিয়ে anonymous struct বানানো যায় 👇

data := struct {
id int
name string
}{
id: 1,
name: "Anonymous",
}

fmt.Println(data)

Output:

{1 Anonymous}

🔹 Struct Inside Struct (Nested Struct)
type Address struct {
city string
zip int
}

type Person struct {
name string
age int
address Address
}

func main() {
p := Person{
name: "Rahim",
age: 25,
address: Address{
city: "Dhaka",
zip: 1207,
},
}

    fmt.Println(p.address.city) // Access nested field

}

🔹 Pointer to Struct

Struct বড় হলে value copy না করে pointer ব্যবহার করা ভালো।

p := Person{name: "Karim", age: 28}
ptr := &p

fmt.Println(ptr.name) // Karim (Go automatically dereferences)
ptr.age = 29
fmt.Println(p.age) // 29

👉 এখানে ptr হলো pointer, কিন্তু Go সহজভাবে ব্যবহার করতে দেয় (তোমাকে \* চিহ্ন দিতে হয় না field access করার সময়)।

🔹 Function with Struct (Method)

Struct এর সাথে method attach করা যায় — অবজেক্ট ওরিয়েন্টেড আচরণ তৈরি করতে।

type Person struct {
name string
age int
}

func (p Person) greet() {
fmt.Println("Hello,", p.name)
}

func main() {
p := Person{name: "Rafiq", age: 22}
p.greet()
}

Output:

Hello, Rafiq

🔹 Pointer Receiver (for modification)

যদি method এর ভিতর struct modify করতে চাও, তাহলে pointer receiver ব্যবহার করো 👇

func (p \*Person) birthday() {
p.age++
}

func main() {
p := Person{name: "Tuhin", age: 24}
p.birthday()
fmt.Println(p.age) // 25
}

🔹 Struct vs Map
Feature Struct Map
Type Safety Strong (fixed types) Dynamic (any key/value)
Compile-time check ✅ Yes ❌ No
Performance Fast Slightly slower
When to use Known fixed fields Dynamic/unpredictable keys
🔹 Real-Life Example (API Response Model)
type User struct {
ID int
Name string
Email string
}

func main() {
user := User{ID: 101, Name: "Sakib", Email: "sakib@gmail.com"}
fmt.Printf("User: %+v\n", user)
}

Output:

User: {ID:101 Name:Sakib Email:sakib@gmail.com}

🧠 মনে রাখার মতো বিষয়
বিষয় মানে
type StructName struct {...} Struct define করা
field access: obj.field Struct field access
method attach: func (r Receiver) method() Struct method
pointer receiver Modify struct
nested struct Struct এর ভিতর struct
anonymous struct একবারের জন্য struct
