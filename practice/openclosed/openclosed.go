package openclosed

type Specification interface {
}

type ColorSpecification struct {
	color Color
}

func (c *ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type Specification interface {
	IsSatisfied(*Product) bool
}

type SizeSpecification struct {
	size Size
}

type AndSpecification struct {
	first, second Specification
}

func (a *AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

func (s *SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type BetterFilter struct {
}

func (b *BetterFilter) Filter(items []Product, spec Specification) []*Product {
	var result []*Product
	for _, i := range items {
		if spec.IsSatisfied(&i) {
			result = append(result, &i)
		}
	}
	return result
}
