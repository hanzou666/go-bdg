package bdg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var createFile func(name string) (io.WriteCloser, error)

func init() {
	createFile = func(name string) (io.WriteCloser, error) {
		return os.Create(name)
	}
}

func (g *Graph) WriteGFA(filename string) error {
	file, err := createFile(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(headerLine()); err != nil {
		return err
	}
	for _, node := range g.Nodes {
		if _, err := writer.WriteString(node.segmentLine()); err != nil {
			return err
		}
	}
	for _, edge := range g.Edges {
		if _, err := writer.WriteString(edge.linkLine()); err != nil {
			return err
		}
	}
	for _, path := range g.Paths {
		if _, err := writer.WriteString(path.pathLine()); err != nil {
			return err
		}
	}
	return writer.Flush()
}

func headerLine() string {
	return "H\tVN:Z:1.0\n"
}

func (n *Node) segmentLine() string {
	return fmt.Sprintf("S\t%d\t%s\n", n.Id, string(n.Seq.Seq))
}

func (e *Edge) linkLine() string {
	return fmt.Sprintf("L\t%d\t%s\t%d\t%s\n", e.FromId, orientStr(e.FromStart), e.ToId, orientStr(e.ToEnd))
}

func orientStr(isReversed bool) string {
	if isReversed {
		return "-"
	} else {
		return "+"
	}
}

func (p *Path) pathLine() string {
	var segmentStr, overlapStr strings.Builder
	for i, mapping := range p.Mappings {
		if i != 0 {
			segmentStr.WriteString(",")
			overlapStr.WriteString(",")
		}
		segmentStr.WriteString(mapping.segmentNameStr())
		overlapStr.WriteString(mapping.overlapStr())
	}
	return fmt.Sprintf("P\t%s\t%s\t%s\n", p.Name, segmentStr.String(), overlapStr.String())
}

func (m *Mapping) segmentNameStr() string {
	return fmt.Sprintf("%d%s", m.Position.NodeId, orientStr(m.Position.IsReversed))
}

func (m *Mapping) overlapStr() string {
	var cigar strings.Builder
	for i, edit := range m.Edits {
		if i != 0 {
			cigar.WriteString(",")
		}
		cigar.WriteString(edit.cigarStr())
	}
	return cigar.String()
}

func (e *Edit) cigarStr() string {
	if e.FromLength == e.ToLength {
		return fmt.Sprintf("%dM", e.FromLength)
	} else {
		if e.FromLength > 0 {
			return fmt.Sprintf("%dD", e.FromLength)
		} else {
			return fmt.Sprintf("%dI", e.ToLength)
		}
	}
}
