import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

def analisar_benchmarks(nome_arquivo_csv) -> None:
    """
    Função para analisar benchmarks de algoritmos a partir de um arquivo CSV.
    Gera gráficos para cada algoritmo mostrando:
    - Tempo de Execução (NsPerOp) vs. Tamanho da Entrada (InputSize)
    - Comparações vs. Tamanho da Entrada
    - Trocas vs. Tamanho da Entrada
    - Memória Alocada (BPerOp) vs. Tamanho da Entrada
    - Quantidade de Alocações (AllocsPerOp) vs. Tamanho da Entrada
    
    """

    # Carregar dados 
    try:
        print("Lendo o arquivo...")
        df = pd.read_csv(nome_arquivo_csv)
        print("Arquivo lido com sucesso.")
    except FileNotFoundError:
        print("Arquivo não encontrado. Certifique-se de que o arquivo existe no caminho especificado.")
        exit(1)

    # Limpando e convertendo dados não numéricos para NaN
    df['Comparisons'] = pd.to_numeric(df['Comparisons'], errors='coerce')
    df['Swaps'] = pd.to_numeric(df['Swaps'], errors='coerce')

    algoritmos = df['Algorithm'].unique()

    for algoritmo in algoritmos:
        print(f"\nGerando gráficos para o algoritmo: {algoritmo}...")
        
        # Filtra o DataFrame para conter apenas os dados do algoritmo atual
        df_algoritmo = df[df['Algorithm'] == algoritmo]

        #Tempo de Execução (NsPerOp) vs. Tamanho da Entrada (InputSize)
        plt.figure(figsize=(12, 8))
        sns.lineplot(data=df_algoritmo, x='InputSize', y='NsPerOp', hue='DataType', marker='o', linestyle='-')
        plt.title(f'Eficiência do Algoritmo: {algoritmo} (Tempo de Execução)')
        plt.xlabel('Tamanho da Entrada (InputSize)')
        plt.ylabel('Tempo Médio (Nanossegundos por Operação)')
        plt.grid(True, which="both", ls="--")
        plt.legend(title='Tipo de Dados')
        nome_arquivo_saida = f'{algoritmo}_tempo_execucao.png'
        plt.savefig(nome_arquivo_saida)
        plt.close()
        print(f" -> Gráfico de tempo de execução salvo como '{nome_arquivo_saida}'")

        #Comparações vs. Tamanho da Entrada
        if not df_algoritmo['Comparisons'].dropna().empty:
            plt.figure(figsize=(12, 8))
            sns.lineplot(data=df_algoritmo, x='InputSize', y='Comparisons', hue='DataType', marker='o', linestyle='-')
            plt.title(f'Eficiência do Algoritmo: {algoritmo} (Comparações)')
            plt.xlabel('Tamanho da Entrada (InputSize)')
            plt.ylabel('Número de Comparações')
            plt.grid(True, which="both", ls="--")
            plt.legend(title='Tipo de Dados')
            nome_arquivo_saida = f'{algoritmo}_comparacoes.png'
            plt.savefig(nome_arquivo_saida)
            plt.close()
            print(f" -> Gráfico de comparações salvo como '{nome_arquivo_saida}'")
        else:
            print(f" -> Não há dados de 'Comparisons' para o algoritmo {algoritmo}.")

        #Trocas vs. Tamanho da Entrada
        if not df_algoritmo['Swaps'].dropna().empty:
            plt.figure(figsize=(12, 8))
            sns.lineplot(data=df_algoritmo, x='InputSize', y='Swaps', hue='DataType', marker='o', linestyle='-')
            plt.title(f'Eficiência do Algoritmo: {algoritmo} (Trocas)')
            plt.xlabel('Tamanho da Entrada (InputSize)')
            plt.ylabel('Número de Trocas (Swaps)')
            plt.grid(True, which="both", ls="--")
            plt.legend(title='Tipo de Dados')
            nome_arquivo_saida = f'{algoritmo}_trocas.png'
            plt.savefig(nome_arquivo_saida)
            plt.close()
            print(f" -> Gráfico de trocas salvo como '{nome_arquivo_saida}'")
        else:
            print(f" -> Não há dados de 'Swaps' para o algoritmo {algoritmo}.")

        #Memória Alocada (BPerOp) vs. Tamanho da Entrada
        if not df_algoritmo['BPerOp'].dropna().empty:
            plt.figure(figsize=(12, 8))
            sns.lineplot(data=df_algoritmo, x='InputSize', y='BPerOp', hue='DataType', marker='o', linestyle='-')
            plt.title(f'Eficiência do Algoritmo: {algoritmo} (Memória Alocada)')
            plt.xlabel('Tamanho da Entrada (InputSize)')
            plt.ylabel('Memória Alocada (Bytes por Operação)')
            plt.grid(True, which="both", ls="--")
            plt.legend(title='Tipo de Dados')
            nome_arquivo_saida = f'{algoritmo}_memoria_alocada.png'
            plt.savefig(nome_arquivo_saida)
            plt.close()
            print(f" -> Gráfico de memória alocada salvo como '{nome_arquivo_saida}'")
        else:
            print(f" -> Não há dados de 'BPerOp' para o algoritmo {algoritmo}.")

        #Quantidade de Alocações (AllocsPerOp) vs. Tamanho da Entrada
        if not df_algoritmo['AllocsPerOp'].dropna().empty:
            plt.figure(figsize=(12, 8))
            sns.lineplot(data=df_algoritmo, x='InputSize', y='AllocsPerOp', hue='DataType', marker='o', linestyle='-')
            plt.title(f'Eficiência do Algoritmo: {algoritmo} (Alocações)')
            plt.xlabel('Tamanho da Entrada (InputSize)')
            plt.ylabel('Número de Alocações por Operação')
            plt.grid(True, which="both", ls="--")
            plt.legend(title='Tipo de Dados')
            nome_arquivo_saida = f'{algoritmo}_alocacoes.png'
            plt.savefig(nome_arquivo_saida)
            plt.close()
            print(f" -> Gráfico de alocações salvo como '{nome_arquivo_saida}'")
        else:
            print(f" -> Não há dados de 'AllocsPerOp' para o algoritmo {algoritmo}.")


nome_do_arquivo_csv = 'benchmark_results.csv'
analisar_benchmarks(nome_do_arquivo_csv)

print("Análise concluída")

