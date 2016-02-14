package mytrie
import (
	"github.com/zhaokunyao/trie"
    "path/filepath"
    "os"
    "os/exec"
)

var tt  = trie.NewTrie()
var t  = trie.NewTrie()

func Init() {
    var err error
    file, _ := exec.LookPath(os.Args[0])
    path, _ := filepath.Abs(file)
    t, err = tt.LoadFromFile(filepath.Dir(path)+"/mytrie/dict")
    if err != nil {
        panic(err)
    }
}

func PrefixMembersList(s string) (members []string) {
    members = t.PrefixMembersList(s)
    return
}

