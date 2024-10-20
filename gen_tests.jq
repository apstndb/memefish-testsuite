.Decls |= (
    map(select(.Name.Name == "TestSQL") |
        .Body.List[] | select(.Lhs[0].Name == "tests") |
            {
                NodeType: "GenDecl", Tok: "var", Specs: [
                {
                    NodeType: "ValueSpec",
                    Names: .Lhs,
                    Type: null,
                    Values: (.Rhs | .[0].Type.Elt.Fields.List |= .[1:] | .[0].Elts[].Elts |= .[1:])
                }
            ]
        }
    )
)
