package twofer

func ShareWith(name string) string {
  if len(name) < 1 {
    name = "you"
  }

  return "One for " + name + ", one for me."
}
