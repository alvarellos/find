# Ctrl + Alt + R ejecuta todo el script

setwd("~/go/src/github.com/estv-admin/find/")

df <- read.csv("values.csv")

library(ggplot2)

ggplot(df, aes(x, y, colour = class)) +
  geom_point()
