// You can edit this code!
// Click here and start typing.
package main

import (
    "fmt"
    "sort"
    "strconv"
    "strings"
)

type KV struct {
    Key   string
    Value int
}
type TreeNode struct {
    Val string
    Len int
    Dir string

    // Map indexed by int. It can be 1 or 0
    // 0 indicates that it's in the left side of the tree.
    // 1 indicates the opposite
    Node map[int]TreeNode
}

func read_node(str string, node *TreeNode, code []string) (string, bool) {
    found := false
    decoded := ""
    for _, n := range node.Node {
        code = append(code, n.Dir)

        if len(n.Node) > 0 {
            decoded, found = read_node(str, &n, code)
            if found {
                break
            } else {
                // remove the last element added
                code = code[:len(code)-1]
            }
        } else {    
            if n.Val == str {
                found = true
                sort.Slice(code, func(i, j int) bool { return code[i] > code[j] })
                decoded = strings.Join(code, "")
                break
            } else {
                // remove the last element added
                code = code[:len(code)-1]
            }
        }
    }
    
    return decoded, found
}

func encode_huff(str string, root *TreeNode) string {
    total := len(str)
    var retval []string

    for i := 0; i < total; i++ {
        str := string(str[i])
        code, _ := read_node(str, root, nil)
        retval = append(retval, code)

    }
    return strings.Join(retval, "")
}

func decode_huff(s string, root *TreeNode) string {    
    total := len(s)
    var word []string
    node := root.Node
    for i := 0; i < total; i++ {
        // get number from char
        str := string(s[i])
        dir, err := strconv.Atoi(str)
        if err != nil {
            return ""
        }

        if len(node[dir].Node) == 0 {
            word = append(word, node[dir].Val)
            node = root.Node
        } else {
            node = node[dir].Node
        }

    }
    return strings.Join(word[:], "")
}

func build_tree(str string) TreeNode {
    total := len(str)
    mapLetterTotal := make(map[string]int)

    for i := 0; i < total; i++ {
        letter := string(str[i])
        if _, ok := mapLetterTotal[letter]; !ok {
            mapLetterTotal[letter] = 1
        } else {
            mapLetterTotal[letter] += 1
        }
    }
    var letters []string
    for letter, _ := range mapLetterTotal {
        letters = append(letters, letter)
    }
    var sorted []KV
    for k, v := range mapLetterTotal {
        sorted = append(sorted, KV{k, v})
    }

    sort.Slice(sorted, func(i, j int) bool {
        return sorted[i].Value < sorted[j].Value
    })

    last := TreeNode{}
    last.Node = make(map[int]TreeNode)

    lastParent := &last
    for i := 0; i < len(sorted); i++ {
        if _, ok := lastParent.Node[1]; !ok {
            lastParent.Len += sorted[i].Value
            lastParent.Node[1] = TreeNode{Val: sorted[i].Key, Len: sorted[i].Value, Dir: "1"}
        } else if _, ok := lastParent.Node[0]; !ok {
            lastParent.Len += sorted[i].Value
            lastParent.Node[0] = TreeNode{Val: sorted[i].Key, Len: sorted[i].Value, Dir: "0"}
        } else {
            parent := TreeNode{Len: lastParent.Len + sorted[i].Value}
            parent.Node = make(map[int]TreeNode)

            lastParent.Dir = "1"
            parent.Node[1] = *lastParent
            parent.Node[0] = TreeNode{Val: sorted[i].Key, Len: sorted[i].Value, Dir: "0"}

            lastParent = &parent
        }

    }

    return *lastParent
}

func main() {
    var input string
    fmt.Scanln(&input)
    tree := build_tree(input)

    code := encode_huff(input, &tree)
    //Enter your code here. Read input from STDIN. Print output to STDOUT

    str := decode_huff(code, &tree)
    fmt.Println(str)
}

