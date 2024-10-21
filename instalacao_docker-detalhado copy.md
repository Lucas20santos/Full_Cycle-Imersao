# Instalacao do docker no linux mint 21.1 e ubuntu 22.04 pelo terminal

Esse passo a passo pode ser encontrado no site do docker

1. Configure apto repositório do Docker.

## Adicione a chave GPG oficial do Docker:

### chaves GPG 

É uma criptografia assimétrica, que ajuda a manter informações longe do alcance de alguns indivíduos.

`

    sudo apt-get update
    sudo apt-get install ca-certificates curl
    sudo install -m 0755 -d /etc/apt/keyrings
    sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
    sudo chmod a+r /etc/apt/keyrings/docker.asc

`

## Adicione o repositório às fontes do Apt:

`

    echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
    $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
    sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

    sudo apt-get update
`

## Comando de instalação

`

     sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

`

## Verificando a instalação

Rode o comando no seu termnal bash:

`

    $ sudo docker run hello-world
`

Mensagem exibida no meu terminal, indicando que o docker foi instalado:

`

    Unable to find image 'hello-world:latest' locally
    latest: Pulling from library/hello-world
    c1ec31eb5944: Pull complete 
    Digest: sha256:d211f485f2dd1dee407a80973c8f129f00d54604d2c90732e8e320e5038a0348
    Status: Downloaded newer image for hello-world:latest

    Hello from Docker!
    This message shows that your installation appears to be working correctly.

    To generate this message, Docker took the following steps:
    1. The Docker client contacted the Docker daemon.
    2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
        (amd64)
    3. The Docker daemon created a new container from that image which runs the
        executable that produces the output you are currently reading.
    4. The Docker daemon streamed that output to the Docker client, which sent it
        to your terminal.

    To try something more ambitious, you can run an Ubuntu container with:
    $ docker run -it ubuntu bash

    Share images, automate workflows, and more with a free Docker ID:
    https://hub.docker.com/

    For more examples and ideas, visit:
    https://docs.docker.com/get-started/

`


## Pode acontecer quando vc tentar atualizar os seus repositorios 

Mesmos os erros abaixo acontecendo, eu consegui instalar o docker normalmente, não precisando substituir por outro repositorio.

`

    Obter:1 https://download.docker.com/linux/ubuntu jammy InRelease [48,8 kB]
    Ign:2 https://mirror.ufscar.br/mint-archive vera InRelease                     
    Atingido:3 https://packages.microsoft.com/repos/code stable InRelease          
    Atingido:4 https://dl.google.com/linux/chrome/deb stable InRelease             
    Atingido:5 https://mirror.ufscar.br/mint-archive vera Release                  
    Atingido:6 http://security.ubuntu.com/ubuntu jammy-security InRelease          
    Obter:7 https://download.docker.com/linux/ubuntu jammy/stable amd64 Packages [40,7 kB]
    Atingido:8 http://repository.spotify.com stable InRelease                      
    Ign:9 https://apt.fury.io/notion-repackaged  InRelease                         
    Ign:10 https://apt.fury.io/notion-repackaged  Release                          
    Ign:11 https://apt.fury.io/notion-repackaged  Packages                         
    Ign:13 https://apt.fury.io/notion-repackaged  Translation-pt                   
    Obter:14 https://packages.konghq.com/public/insomnia/deb/ubuntu focal InRelease [2.942 B]
    Ign:15 https://apt.fury.io/notion-repackaged  Translation-en                   
    Ign:16 https://apt.fury.io/notion-repackaged  Translation-pt_BR                
    Obter:11 https://apt.fury.io/notion-repackaged  Packages [1.572 B]             
    Ign:13 https://apt.fury.io/notion-repackaged  Translation-pt                   
    Ign:15 https://apt.fury.io/notion-repackaged  Translation-en                   
    Ign:16 https://apt.fury.io/notion-repackaged  Translation-pt_BR                
    Ign:13 https://apt.fury.io/notion-repackaged  Translation-pt                   
    Ign:15 https://apt.fury.io/notion-repackaged  Translation-en                   
    Ign:16 https://apt.fury.io/notion-repackaged  Translation-pt_BR                
    Ign:13 https://apt.fury.io/notion-repackaged  Translation-pt                   
    Ign:15 https://apt.fury.io/notion-repackaged  Translation-en                   
    Ign:16 https://apt.fury.io/notion-repackaged  Translation-pt_BR                
    Ign:13 https://apt.fury.io/notion-repackaged  Translation-pt                   
    Ign:15 https://apt.fury.io/notion-repackaged  Translation-en                   
    Ign:16 https://apt.fury.io/notion-repackaged  Translation-pt_BR                
    Ign:13 https://apt.fury.io/notion-repackaged  Translation-pt
    Ign:15 https://apt.fury.io/notion-repackaged  Translation-en
    Ign:16 https://apt.fury.io/notion-repackaged  Translation-pt_BR
    Ign:13 https://apt.fury.io/notion-repackaged  Translation-pt                   
    Ign:15 https://apt.fury.io/notion-repackaged  Translation-en                   
    Ign:16 https://apt.fury.io/notion-repackaged  Translation-pt_BR                
    Atingido:17 https://deb.opera.com/opera-stable stable InRelease                
    Ign:18 http://mirror.unesp.br/ubuntu jammy InRelease                           
    Ign:19 http://mirror.unesp.br/ubuntu jammy-updates InRelease
    Ign:20 http://mirror.unesp.br/ubuntu jammy-backports InRelease
    Ign:18 http://mirror.unesp.br/ubuntu jammy InRelease
    Ign:19 http://mirror.unesp.br/ubuntu jammy-updates InRelease
    Ign:20 http://mirror.unesp.br/ubuntu jammy-backports InRelease
    Ign:18 http://mirror.unesp.br/ubuntu jammy InRelease
    Ign:19 http://mirror.unesp.br/ubuntu jammy-updates InRelease
    Ign:20 http://mirror.unesp.br/ubuntu jammy-backports InRelease
    Err:18 http://mirror.unesp.br/ubuntu jammy InRelease
    Não foi possível iniciar a conexão para mirror.unesp.br:80 (2801:88:1000::5). - connect (101: A rede está fora de alcance) Não foi possível conectar em mirror.unesp.br:80 (200.145.6.5), conexão expirou
    Err:19 http://mirror.unesp.br/ubuntu jammy-updates InRelease
    Não foi possível iniciar a conexão para mirror.unesp.br:80 (2801:88:1000::5). - connect (101: A rede está fora de alcance)
    Err:20 http://mirror.unesp.br/ubuntu jammy-backports InRelease
    Não foi possível iniciar a conexão para mirror.unesp.br:80 (2801:88:1000::5). - connect (101: A rede está fora de alcance)
    Baixados 94,0 kB em 39s (2.388 B/s)
    Lendo listas de pacotes... Pronto
    W: Falhou ao buscar http://mirror.unesp.br/ubuntu/dists/jammy/InRelease  Não foi possível iniciar a conexão para mirror.unesp.br:80 (2801:88:1000::5). - connect (101: A rede está fora de alcance) Não foi possível conectar em mirror.unesp.br:80 (200.145.6.5), conexão expirou
    W: Falhou ao buscar http://mirror.unesp.br/ubuntu/dists/jammy-updates/InRelease  Não foi possível iniciar a conexão para mirror.unesp.br:80 (2801:88:1000::5). - connect (101: A rede está fora de alcance)
    W: Falhou ao buscar http://mirror.unesp.br/ubuntu/dists/jammy-backports/InRelease  Não foi possível iniciar a conexão para mirror.unesp.br:80 (2801:88:1000::5). - connect (101: A rede está fora de alcance)
    W: Falhou o download de alguns ficheiros de índice. Foram ignorados ou os antigos foram usados em seu lugar
`

O "erro" que você está enfrentando ao tentar baixar pacotes do repositório da UNESP (mirror.unesp.br) indica que o servidor não está acessível. Isso pode ser causado por problemas de rede ou indisponibilidade do servidor.

Só use esse comando abaixo, caso vc não consiga instalar, o usando oS comandoS de instalação do docker. Esses comandos estarão logo acima.

Aqui estão algumas opções para resolver o problema:

1. Alterar o mirror:

- Você pode alterar o espelho do repositório para um que esteja funcionando. Para isso, edite o arquivo /etc/apt/sources.list e substitua as linhas relacionadas ao mirror.unesp.br por outro espelho, como o mirror oficial da Canonical.

Exemplo: 

`
    
    abra o terminal bash e digite: 

    $ sudo sed -i 's|mirror.unesp.br|archive.ubuntu.com|g' /etc/apt/sources.list

`

2. Verificar a conectividade da rede:

Certifique-se de que sua conexão de rede está funcionando corretamente e que você consegue acessar outros sites.

3. Atualizar o APT sem o mirror da UNESP:

Se você quiser ignorar temporariamente o mirror da UNESP, você pode comentar ou remover as linhas que mencionam o mirror.unesp.br no arquivo /etc/apt/sources.list.

Execute o seguinte comando para abrir o arquivo:

`
    No terminal bash e digite: 
    
    sudo nano /etc/apt/sources.list

`
E comente ou remova as linhas que mencionam o espelho da UNESP

Após essas alterações, você pode tentar atualizar os pacotes novamente com:

`
    No terminal bash e digite: 
    
    sudo apt update

`

## Descrição de cada comando de instalação

#### sudo apt-get update

Atualiza a lista de pacotes disponíveis no gerenciador de pacotes do Ubuntu (APT). Isso garante que o sistema saiba quais pacotes 
estão disponíveis para instalação e suas versões mais recentes.

#### sudo apt-get install ca-certificates curl

Esse comando instala dois pacotes

- ca-certificates: Um pacote que instala certificados digitais de autoridades certificadoras (CAs)
- curl: Uma ferramenta de linha de comando usada para transferir dados via URLs.

#### sudo install -m 0755 -d /etc/apt/keyrings
 
Neste contexto, o comando install (não confudir com o apt install) é usado para criar diretórios ou copiar arquivos com permissões 
específicas

- -m 0755: Define as permissões do diretório criado. O 0755 significa:
    - O dono do diretório tem permissões de leitura, gravação e execução (7)
    - Outros usuários têm permissões de leitura e execução(%)
    - Outros usuários têm permissões de leitura e execução(5)
- -d: Indica que o objetivo é criar um diretório.
- /etc/apt/keyrings: O caminho onde o diretório será criado. Esse diretório será usado para armazenar chaves de assinatura GPG.

#### sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc

- curl: ferramenta que baixa conteúdo da internet usando uma URL.
- -fsSL: conjunto de opções do curl:
    - -f: Fails silently(não mostra mensagens de erro em caso de falha).
    - -s: Silent mode (não exibe barra de progresso ou outras mensagens desnecessárias)
    - -S: Mostra mensagens de erro se houverem.
    - -L: Segue redirecionamentos, útil se o URL para redirecionado para outro local.
- https://download.docker.com/linux/ubuntu/gpg: O URL de onde a chave GPG do Docker será baixada. Essa chave é usada para verificar a autenticidade dos pacotes Docker
- -o /etc/apt/keyrings/docker.asc: Salva o arquivo baixado (neste caso, a chave GPG) no caminho -o /etc/apt/keyrings/docker.asc.
- sudo chmod a+r /etc/apt/keyrings/docker.asc: 

#### sudo chmod a+r /etc/apt/keyrings/docker.asc

- chmod: Comando que modifica as permissões de arquivos ou diretórios.
- a+r: Atribui a todos os usuários (a) permissão de leitura (r) para o arquivo.
- /etc/apt/keyrings/docker.asc: Especifica o arquivo para o qual as permissões estão sendo ajustadas. No caso, o arquivo da chave 
GPG do Docker.

#### echo \  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \   $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \ sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

- echo: Imprime uma linha de texto, que é o repositório Docker a ser adicionado
- [arch=$(dpkg --print-architecture)]: Insere a arquitetura do sistema (por exemplo, amd64) no arquivo de fontes do Docker. 
    - O dpkg --print-architecture retorna a arquitetura da máquina, como amd64 ou arm64.
- signed-by=/etc/apt/keyrings/docker.asc: Informa ao APT para usar a chave GPG (para assinar pacotes) localizada em /etc/apt/keyrings/docker.asc.
- https://download.docker.com/linux/ubuntu: URL do repositório do Docker para o Ubuntu.
- $(. /etc/os-release && echo "$UBUNTU_CODENAME"): Pega o codename da sua versão do linux-mint a partir do arquivo /etc/os-release e
    - UBUNTU_CODENAME: Caso for o Ubuntu
    - VERSION_CODENAME: Substituir caso seja um derivado do Ubuntu, como nesse caso o linux-mint

#### sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

- sudo apt-get install + 
    - docker-ce: Este é o Docker Community Edition, ou seja, a versão gratuita e de código aberto do Docker, destinada a desenvolvedores e usuários que querem usar contêineres em seus sistemas. O Docker CE inclui o Docker Engine, que é o núcleo do Docker e permite criar, gerenciar e executar contêineres.
    - docker-ce-cli: O Docker CLI (Interface de Linha de Comando) é o programa de linha de comando que você usa para interagir com o Docker. Ele permite que você execute comandos como: 
        - docker run, 
        - docker pull, 
        - docker build, 
        - e muitos outros. 
    Ele se comunica com o Docker Engine para gerenciar contêineres, redes e imagens.
    - containerd.io:
        - O containerd é o runtime de contêiner de nível mais baixo que o Docker usa para gerenciar o ciclo de vida dos contêineres. É uma camada importante entre o Docker Engine e os sistemas operacionais que o Docker gerencia.
        - containerd.io é a implementação do containerd que gerencia as operações fundamentais de contêineres, como transferências de imagem, armazenamento de imagens, execução e controle de contêineres.
    - docker-buildx-plugin:
        - O Docker Buildx é uma extensão do Docker CLI que fornece recursos avançados de construção de imagens de contêiner. Ele suporta a criação de imagens multi-plataforma, o que significa que você pode construir uma única imagem de contêiner que funcione em diferentes arquiteturas, como ARM e x86.
        - Este plugin permite a execução de comandos avançados de construção, como docker buildx build.
    - docker-compose-plugin:
        - O Docker Compose Plugin adiciona suporte ao Docker Compose, uma ferramenta que permite definir e gerenciar aplicações que envolvem múltiplos contêineres. O Compose utiliza um arquivo YAML (docker-compose.yml) para definir a configuração dos serviços, e com um único comando, você pode orquestrar todos os contêineres da aplicação.
        - Ele é extremamente útil para aplicações que têm dependências em vários contêineres (por exemplo, um contêiner para um banco de dados, outro para um servidor web, etc.).

#### Como tudo funciona junto:

- Docker Engine (docker-ce) é o núcleo que lida com a execução e gerenciamento de contêineres.
- Docker CLI (docker-ce-cli) permite que você interaja com o Docker através de comandos.
- containerd.io gerencia o ciclo de vida dos contêineres em segundo plano.
- Buildx Plugin expande as capacidades de construção de imagens de contêiner.
- Compose Plugin permite a orquestração e gerenciamento de múltiplos contêineres em uma única aplicação.

Esses pacotes juntos permitem que você crie, gerencie, execute e orquestre contêineres no Docker com todas as funcionalidades disponíveis para desenvolvedores e administradores de sistema.
