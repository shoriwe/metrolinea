import json
import os.path
import re
import sys

import requests


def load() -> dict:
	if not os.path.exists('metrolinea.json'):
		with open("metrolinea.json", "wb") as file:
			response = requests.get("https://www.datos.gov.co/resource/kcdt-jbvj.json")
			file.write(response.content)
	with open("metrolinea.json", "rb") as file:
		return json.load(file)


def generate_graph(data: dict) -> list:
	terminales = {}
	puntos_terminales = {}
	for entry in data:
		try:
			codigo = entry["codigo"]
			# ruta = tuple(re.split(r"\s*[-–]\s*", entry["ruta"]))
			# ruta = tuple(re.split(r"\s*[-–]\s*", entry["recorrido"]))
			ruta = tuple(re.split(r"\s*[-–]\s*", entry["cartel_de_ruta_ida"]))
			distancia = float(
				re.search(r"\d+\.{0,1}\d*", entry.setdefault("long_km", str(sys.maxsize)))
					.group(0))
		except KeyError:
			continue
		puntos_terminales[ruta[0]] = entry["terminal"]
		try:
			terminales.setdefault(entry["terminal"], {}).setdefault(
				ruta[0], {})[ruta] = {
				"codigo": codigo,
				"destino": ruta[-1],
				"distancia-km": distancia,
				"minutos-valle:": float(entry["frecuencia_despacho_hora_valle"]),
				"minutos-pico:": float(entry["frecuencia_de_despacho_hora_pico"])
			}
		except KeyError as e:
			pass
	print(terminales)
	for terminales in terminales.values():
		for punto in terminales.values():
			for ruta in punto.values():
				try:
					terminal = puntos_terminales[ruta["destino"]]
					print(terminal)
				except KeyError:
					print(data[int(ruta["codigo"])])
					print(ruta)
					pass
	graph = {}
	return graph


def main():
	data = load()
	graph = generate_graph(data)
	with open("metrolinea-graph.json", "w") as file:
		json.dump(graph, file)


if __name__ == "__main__":
	main()
