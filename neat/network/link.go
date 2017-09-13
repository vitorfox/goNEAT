package network

import (
	"github.com/yaricom/goNEAT/neat"
	"fmt"
	"flag"
)

// A LINK is a connection from one node to another with an associated weight.
// It can be marked as recurrent.
type Link interface {

	// Set added weight
	SetAddedWeight(weight float64)
	// Returns IN node
	InNode() *NNode
	// Returns link weight
	GetWeight() float64

	// Returns true if link is time delayed
	IsTimeDelayed() bool
	// Returns true if link is recurrent
	IsRecurrent() bool
}

// Creates new link with specified weight, input and output neurons connected reccurently or not.
func NewLink(weight float64, innode, outnode *NNode, recurrent bool) Link {
	link := newLink(weight)
	link.in_node = innode
	link.out_node = outnode
	link.is_recurrent = recurrent
	return link
}

// Creates new Link with specified Trait
func NewLinkWithTrait(trait *Trait, weight float64, innode, outnode *NNode, recurrent bool) Link {
	link := newLink(weight)
	link.in_node = innode
	link.out_node = outnode
	link.is_recurrent = recurrent
	link.linktrait = trait
	return link
}

func NewLinkWeight(weight float64) Link {
	return newLink(weight)
}

// The internal representation
type link struct {
	// Weight of connection
	weight float64
	// NNode inputting into the link
	in_node *NNode
	// NNode that the link affects
	out_node *NNode
	// If TRUE the link is recurrent
	is_recurrent bool
	// If TRUE the link is time delayed
	time_delay bool

	// Points to a trait of parameters for genetic creation
	linktrait *Trait

	/* ************ LEARNING PARAMETERS *********** */
	/* These are link-related parameters that change
	   during Hebbian type learning */
	// The amount of weight adjustment
	added_weight float64
	// The parameters to be learned
	params []float64
}

// The private default constructor
func newLink(weight float64) link {
	return link{
		weight:weight,
		params:make([]float64, neat.Num_trait_params),
	}
}

// The Link interface implementation
func (l *link) SetAddedWeight(weight float64) {
	l.added_weight = weight
}
func (l *link) InNode() *NNode {
	return l.in_node
}
func (l *link) GetWeight() float64 {
	return l.weight
}
func (l *link) IsTimeDelayed() bool {
	return l.time_delay
}
func (l *link) IsRecurrent() bool {
	return l.is_recurrent
}

func (n link) String() string {
	return fmt.Sprintf("(link: %s <-> %s, w: %f", n.in_node, n.out_node, n.weight)
}