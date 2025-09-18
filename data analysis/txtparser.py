import re
import csv

def parse_benchmark_data(file_path='benchmark.txt'):
    """
    Analisa um arquivo de log de benchmark Go, extrai dados de desempenho e estatísticas
    e os combina em uma única estrutura de dados.

    Args:
        file_path (str): O caminho para o arquivo de benchmark.

    Returns:
        list: Uma lista de dicionários, onde cada dicionário representa uma linha
              de dados combinados de benchmark.
    """
   
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            lines = f.readlines()
    except FileNotFoundError:
        print(f"Erro: O arquivo '{file_path}' não foi encontrado.")
        return []

    statistics = {}
    # Regex para extrair estatísticas de comparações e trocas
    stats_header_re = re.compile(r'==\s*(?P<algorithm>\w+)\s*==')
    stats_type_re = re.compile(r'--\s*(?P<data_type>\w+)\s*--')
    stats_line_re = re.compile(r'\s*(?P<size>\d+):\s*Comparisons:\s*(?P<comparisons>\d+),\s*Swaps:\s*(?P<swaps>\d+)')

    current_algorithm = None
    current_data_type = None

    #Extrair todas as estatísticas (comparações e trocas)
    for line in lines:
        line_clean = re.sub(r'\\s*', '', line).strip()
        
        header_match = stats_header_re.match(line_clean)
        if header_match:
            current_algorithm = header_match.group('algorithm')
            continue

        type_match = stats_type_re.match(line_clean)
        if type_match:
            current_data_type = type_match.group('data_type')
            continue

        stats_match = stats_line_re.match(line_clean)
        if stats_match and current_algorithm and current_data_type:
            size = int(stats_match.group('size'))
            comparisons = int(stats_match.group('comparisons'))
            swaps = int(stats_match.group('swaps'))
            key = (current_algorithm, current_data_type, size)
            statistics[key] = {'Comparisons': comparisons, 'Swaps': swaps}

    #Extrair resultados de benchmark combinando com as estatísticas
    combined_data = []
    full_text = "".join(lines)

    # Regex para encontrar todos os blocos de resultados de benchmark
    benchmark_results_re = re.compile(
        r'(BenchmarkSortingAlgorithms/(?P<algorithm>\w+)-(?P<data_type>\w+)-(?P<size>\d+)-\d+\s+'
        r'(?P<iterations>\d+)\s+'
        r'(?P<ns_per_op>[\d\.]+)\s*ns/op\s+'
        r'(?P<b_per_op>\d+)\s*B/op\s+'
        r'(?P<allocs_per_op>\d+)\s*allocs/op)',
        re.MULTILINE
    )

    # Limpar todo o texto de uma vez para facilitar a correspondência 
    full_text_clean = re.sub(r'\\s*', '', full_text)
    
    for match in benchmark_results_re.finditer(full_text_clean):
        algorithm = match.group('algorithm')
        data_type = match.group('data_type')
        size = int(match.group('size'))
        
        stats_key = (algorithm, data_type, size)
        stats = statistics.get(stats_key, {'Comparisons': 'N/A', 'Swaps': 'N/A'})

        combined_data.append({
            'Algorithm': algorithm,
            'DataType': data_type,
            'InputSize': size,
            'Iterations': int(match.group('iterations')),
            'NsPerOp': float(match.group('ns_per_op')),
            'BPerOp': int(match.group('b_per_op')),
            'AllocsPerOp': int(match.group('allocs_per_op')),
            'Comparisons': stats['Comparisons'],
            'Swaps': stats['Swaps']
        })
        
    return combined_data

def write_to_csv(data, filename='benchmark_results.csv'):

    if not data:
        print("Nenhum dado para escrever no CSV.")
        return

    # Ordenar os dados para uma saída consistente
    data.sort(key=lambda x: (x['Algorithm'], x['DataType'], x['InputSize']))

    fieldnames = ['Algorithm', 'DataType', 'InputSize', 'Iterations', 'NsPerOp', 
                  'BPerOp', 'AllocsPerOp', 'Comparisons', 'Swaps']
    
    with open(filename, 'w', newline='', encoding='utf-8') as csvfile:
        writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
        writer.writeheader()
        writer.writerows(data)
    
    print(f"Dados de benchmark foram processados e salvos com sucesso em '{filename}'")

if __name__ == "__main__":
    parsed_data = parse_benchmark_data('Caminho/para/o/arquivo/benchmark.txt')
    if parsed_data:
        write_to_csv(parsed_data)