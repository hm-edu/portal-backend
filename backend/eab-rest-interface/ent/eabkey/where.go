// Code generated by entc, DO NOT EDIT.

package eabkey

import (
	"entgo.io/ent/dialect/sql"
	"github.com/hm-edu/eab-rest-interface/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// User applies equality check predicate on the "user" field. It's identical to UserEQ.
func User(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUser), v))
	})
}

// EabKey applies equality check predicate on the "eabKey" field. It's identical to EabKeyEQ.
func EabKey(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEabKey), v))
	})
}

// UserEQ applies the EQ predicate on the "user" field.
func UserEQ(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUser), v))
	})
}

// UserNEQ applies the NEQ predicate on the "user" field.
func UserNEQ(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUser), v))
	})
}

// UserIn applies the In predicate on the "user" field.
func UserIn(vs ...string) predicate.EABKey {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EABKey(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUser), v...))
	})
}

// UserNotIn applies the NotIn predicate on the "user" field.
func UserNotIn(vs ...string) predicate.EABKey {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EABKey(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUser), v...))
	})
}

// UserGT applies the GT predicate on the "user" field.
func UserGT(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUser), v))
	})
}

// UserGTE applies the GTE predicate on the "user" field.
func UserGTE(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUser), v))
	})
}

// UserLT applies the LT predicate on the "user" field.
func UserLT(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUser), v))
	})
}

// UserLTE applies the LTE predicate on the "user" field.
func UserLTE(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUser), v))
	})
}

// UserContains applies the Contains predicate on the "user" field.
func UserContains(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUser), v))
	})
}

// UserHasPrefix applies the HasPrefix predicate on the "user" field.
func UserHasPrefix(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUser), v))
	})
}

// UserHasSuffix applies the HasSuffix predicate on the "user" field.
func UserHasSuffix(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUser), v))
	})
}

// UserEqualFold applies the EqualFold predicate on the "user" field.
func UserEqualFold(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUser), v))
	})
}

// UserContainsFold applies the ContainsFold predicate on the "user" field.
func UserContainsFold(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUser), v))
	})
}

// EabKeyEQ applies the EQ predicate on the "eabKey" field.
func EabKeyEQ(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEabKey), v))
	})
}

// EabKeyNEQ applies the NEQ predicate on the "eabKey" field.
func EabKeyNEQ(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEabKey), v))
	})
}

// EabKeyIn applies the In predicate on the "eabKey" field.
func EabKeyIn(vs ...string) predicate.EABKey {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EABKey(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEabKey), v...))
	})
}

// EabKeyNotIn applies the NotIn predicate on the "eabKey" field.
func EabKeyNotIn(vs ...string) predicate.EABKey {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EABKey(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEabKey), v...))
	})
}

// EabKeyGT applies the GT predicate on the "eabKey" field.
func EabKeyGT(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEabKey), v))
	})
}

// EabKeyGTE applies the GTE predicate on the "eabKey" field.
func EabKeyGTE(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEabKey), v))
	})
}

// EabKeyLT applies the LT predicate on the "eabKey" field.
func EabKeyLT(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEabKey), v))
	})
}

// EabKeyLTE applies the LTE predicate on the "eabKey" field.
func EabKeyLTE(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEabKey), v))
	})
}

// EabKeyContains applies the Contains predicate on the "eabKey" field.
func EabKeyContains(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEabKey), v))
	})
}

// EabKeyHasPrefix applies the HasPrefix predicate on the "eabKey" field.
func EabKeyHasPrefix(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEabKey), v))
	})
}

// EabKeyHasSuffix applies the HasSuffix predicate on the "eabKey" field.
func EabKeyHasSuffix(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEabKey), v))
	})
}

// EabKeyEqualFold applies the EqualFold predicate on the "eabKey" field.
func EabKeyEqualFold(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEabKey), v))
	})
}

// EabKeyContainsFold applies the ContainsFold predicate on the "eabKey" field.
func EabKeyContainsFold(v string) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEabKey), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EABKey) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EABKey) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EABKey) predicate.EABKey {
	return predicate.EABKey(func(s *sql.Selector) {
		p(s.Not())
	})
}