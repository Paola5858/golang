import math
from turtle import *

# Funções para calcular as coordenadas do coração
def hearta(k):
    return 15 * math.sin(k) ** 3

def heartb(k):
    return 12 * math.cos(k) - 5 * math.cos(2 * k) - 2 * math.cos(3 * k) - math.cos(4 * k)

# Configuração do desenho
speed(0)
bgcolor("black")
penup()

title("Coração Interativo")

# Desenhar o coração com efeito mais suave
for i in range(10000):  # Mais pontos para suavizar o traço
    x = hearta(i * 0.02) * 20  # Ajuste de escala
    y = heartb(i * 0.02) * 20
    goto(x, y)
    pendown()
    color("red")
    dot(2)  # Pontos menores para efeito mais fluido
    penup()

done()

