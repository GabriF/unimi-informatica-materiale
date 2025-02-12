---
title: Fisica
date: a.a. 2024-2025
author: Gabriele Fioco
lang: it-IT
geometry: margin=1.5in
toc: true
toc-depth: 2
numbersections: true
colorlinks: true
linkcolor: black
urlcolor: black 
toccolor: black
header-includes: |
    \usepackage{fancyhdr}
    \usepackage{xcolor}
    \usepackage{mathtools}
    \usepackage{tikz}
    \usepackage{amsmath}
    \usepackage{float}
    \pagestyle{fancy}
...

\graphicspath{ {img/} }
\counterwithin{figure}{section}
\counterwithin{equation}{section}

# Cinematica

Un **punto materiale** è un oggetto che ha una massa ma non ha dimensioni, la cui posizione è descritta da un vettore in $\mathbb{R}^3$.

## Traiettoria

La **traiettoria** è il percorso che compie un punto materiale nel tempo, descritta dalla legge oraria nella funzione \ref{eq:legge_oraria}

\begin{equation} \label{eq:legge_oraria}
\vec{x} : [t_0, t] \rightarrow \mathbb{R}^3
\end{equation}

La funzione \ref{eq:legge_oraria} vale finchè il punto non subisce interazioni con altri corpi. Il dominio della funzione $[t_0, t]$ è l'intervallo di tempo in cui si osserva il movimento, il codominio è lo spazio tridimensionale rappresentato da $\mathbb{R}^3$, che indicheremo con le coordinate $(x, y, z) = \hat i x + \hat j y + \hat k z$.

Lo spazio percorso lungo la traiettoria è una grandezza fisica misurata, espressa dimensionalmente come una lunghezza $[L]$ in unità fisiche (es. metri).

## Velocità e accelerazione

La **velocità** descrive quanto rapidamente varia la misura di un punto in funzione del tempo. Considerato un punto che si muove tra due punti $x(t_2)$ e $x(t_1)$ in un tempo $t_2 - t_1$ con $t_2 > t_1$ la velocità media è così data

\begin{equation} \label{eq:vel_media}
\hat v_{\text{media}} = \frac{x(t_2) - x(t_1)}{t_2 - t_1} = \frac{\Delta \vec x}{\Delta t}
\end{equation}

Geometricamente la velocità è il rapporto incrementale della traiettoria. Siamo interessati al calcolo della **velocità istantanea** che si ottiene quando $\delta \hat x \rightarrow 0$ e $\delta t \rightarrow 0$. Quindi la velocità istantanea è la derivata della traiettoria rispetto al tempo.

\begin{equation} \label{eq:velocita}
\vec v(t) = \lim_{\Delta t \rightarrow 0} \frac{\Delta \vec x}{\Delta t} = \frac{d}{dt} \vec x(t) = \frac{d}{dt} (x, y, z) = \frac{d}{dt} (\hat i x + \hat j y + \hat k z) = \frac{d}{dt} \hat i x + \frac{d}{dt} \hat j y + \frac{d}{dt} \hat k z
\end{equation}

Dimensionalmente la velocità è $[v] = \frac{[L]}{[T]}$ dove $L$ è una lunghezza e $T$ un tempo.

Similmente l'**accelerazione media** è data da

\begin{equation} \label{eq:accel_media}
\vec a = \frac{v(t_2) - v(t_1)}{t_2 - t_1}
\end{equation}

L'**accelerazione istantanea** è la derivata rispetto alla velocità

\begin{equation} \label{eq:accelerazione}
a(t) = \lim_{\Delta t \rightarrow 0} \frac{\Delta \vec v}{\Delta t} = \frac{d^2 x}{d t}
\end{equation}

Dimensionalmente l'accelerazione è $[a] = \frac{[L]}{[T^2]}$.

## Moto rettilineo uniforme e uniformemente accelerato

Data un'accelerazione generica $a$ ci si chiede quale posizione $x$ assume al tempo $t$, quindi si vuole la traiettoria. Ricordando che l'accelerazione è la derivata della velocità, che è a sua volta la derivata della traiettoria, si può integrare due volte rispetto a $t$. Si ottiene l'equazione \ref{eq:mua}

\begin{equation} \label{eq:mua}
\vec x(t) = \frac{1}{2} a t^2 + bt + c
\end{equation}

Si osserva che $b$ e $c$ rappresentano rispettivamente la velocità iniziale e la posizione iniziale che indicheremo con $\vec v_0$ e $\vec x_0$. L'equazione \ref{eq:mua} si riscrive

\begin{equation} \label{eq:mua_en}
\vec x(t) = \frac{1}{2} \vec a t^2 + \vec v_0 t + \vec x_0
\end{equation}

Ora ci si chiede se si può esprimere questo moto senza usare il tempo. Derivando l'equazione \ref{eq:mua_en} si ottiene $\vec v(t) = \vec a t + \vec v_0$, quindi

\begin{equation} \label{eq:formula_tempo_mua}
t = \frac{\vec v(t) - \vec v_0}{a}
\end{equation}

Sostituendo la \ref{eq:formula_tempo_mua} nella \ref{eq:mua_en} si ottiene la relazione tra posizione, velocità e accelerazione senza riferimenti al tempo

\begin{equation} \label{eq:formula_mua}
2 \vec a(\vec x(t) - \vec x_0) = \vec v(t)^2 - \vec v_0^2
\end{equation}

## Moto circolare uniforme

Il moto circolare uniforme è il moto di un punto che si muove su una traiettoria di raggio $R$ con velocità costante.

\begin{figure}[ht]
\includegraphics[width=0.35\textwidth]{moto_circolare}
\centering
\caption{Moto circolare uniforme}
\end{figure}

Su un piano a due dimensioni la traiettoria è definita, nelle sue componenti, come:

\begin{equation} \label{eq:moto_circolare_comp}
\begin{cases}
x(t) = R \cos (\omega t) \\
y(t) = R \sin (\omega t)
\end{cases}
\end{equation}

In forma vettoriale:

\begin{equation} \label{eq:moto_circolare}
\vec x(t) = \hat i R \cos (\omega t) + \hat j \sin (\omega t)
\end{equation}

Si osserva che il punto ritorna al punto di inizio in un tempo $T$, detto periodo, quando l'angolo $\omega T = 2 \pi$, dunque $\omega = \frac{2 \pi}{T} = 2 \pi f$ dove $f = \frac{1}{T}$ è detta frequenza. Il periodo rappresenta il tempo per fare un giro, la frequenza il numero di giri effettuati in un secondo, $\omega$ definita come $\omega = \frac{\theta}{t}$ è detta velocità angolare e rappresenta i radianti passati al secondo.

Derivando la \ref{eq:moto_circolare} rispetto a $t$ si ottiene la velocità:

\begin{equation} \label{eq:moto_circolare_v}
\vec v(t) = - \hat i \omega R \sin(\omega t) + \hat j \omega \cos (\omega t)
\end{equation}

Calcolando il modulo della velocità:

\begin{equation}
|\vec v| = v = \sqrt{V_x^2 + V_y^2} = \omega R
\end{equation}

$v$ è detta velocità tangenziale ed è sempre tangente alla circonferenza.

Derivando la \ref{eq:moto_circolare_v} si ottiene l'accelerazione:

\begin{equation} \label{eq:moto_circolare_a}
\vec a(t) = - \omega^2 R \hat i \cos (\omega t) - \omega^2 R \hat j \sin (\omega t) = - \omega^2 R (\hat i \cos(\omega t) + \hat j \sin (\omega t)) = - \omega^2 R
\end{equation}

Calcolando il modulo dell'accelerazione:

\begin{equation}
|a| = a = \omega^2 R = \frac{v^2}{R}
\end{equation}

$a$ è detta accelerazione centripeta ed indica che il punto è diretto verso il centro.

## Moto parabolico

Il **moto parabolico** è un moto in uno spazio 2D, rappresentante un proiettile, composto da due componenti: un moto rettilineo in orizzontale e un moto rettilineo uniformemente accelerato in verticale. Quest'ultimo tipicamente utilizza l'accelerazione gravitazionale, indicata con $g = -9,8 m/s^2$.

\begin{figure}[ht]
\includegraphics[width=0.35\textwidth]{moto_parabolico}
\centering
\caption{Moto parabolico, $R$ è la gittata e $y_{max}$ l'altezza massima}
\end{figure}

Le due componenti sono:

\begin{equation} \label{eq:moto_parabolico}
\begin{cases}
x(t) = v_0\cos(\Theta)t \\
y(t) = v_0\sin(\Theta)t - \frac{1}{2}gt^2
\end{cases}
\end{equation}

Derivando si ottiene la **velocità**:

\begin{equation}
\begin{cases}
v_x(t) = v_0\cos(\Theta) \\
v_y(t) = v_0\sin(\Theta)t - gt
\end{cases}
\end{equation}

L'**accelerazione** è fissata a $g$, come da definizione.

Risolvendo la \ref{eq:moto_parabolico} per $t$ si ottiene:

\begin{equation} \label{eq:parabolico_tempo}
t = \frac{x}{v_0 \cos\theta}
\end{equation}

Sostituendo in $y(t)$:

\begin{equation}
y(x) = y_0 + x\tan\theta - \frac{g}{2(v_0\cos\theta)^2}x^2
\end{equation}

Tempo totale
: Il tempo totale $t_f$ si ottiene quando $y(t) = 0$, si calcola:

\begin{equation}
t_f = \frac{2 v_0 \sin \theta}{g}
\end{equation}

Gittata
: Per trovare la gittata $x_f$ si calcola la componente orizzontale al tempo totale, cioè $x(t_f)$ e si ottiene:

Si osservano diversi comportamenti nella gittata:

- La gittata massima si ottiene quando $sen(2\theta) = 1$ ovvero quando $2\theta = 90 \deg$ ovvero quando $\theta = 45 \deg$.
- Per $\theta < 45 \deg$ allora la gittata diminuisce.
- Per $\theta > 45 \deg$ la gittata diminuisce, aumenta la velocità verticale ma diminuisce quella orizzontale.
- Il moto parabolico è simmetrico, quindi angoli simmetrici hanno stessa gittata.

\begin{equation}
x_f = \frac{v_0^2 \sin(2\theta)}{g}
\end{equation}

Altezza massima
: L'altezza massima $(x_M, y_M)$ si raggiunge quando $t = t_f / 2 = t_M$, a causa della simmetria del moto parabolico. Quindi si calcolano le componenti come $x(t_M)$ e $y(t_M)$ e si ottiene: 

\begin{equation}
x_M = \frac{v_0^2\sin(2\theta)}{2g} v_0^2
\end{equation}

\begin{equation}
y_M = \frac{v_0^2 \sin^2(2\theta)}{2g} v_0^2
\end{equation}

L'altezza massima si può ottenere anche imponendo la componente della velocità verticale pari a 0.

Velocità all'impatto
: La velocità all'impatto rimane $v_0 \cos\theta$ nella componente orizzontale, mentre è $v_x(t_f)$ nella sua componente verticale:

\begin{equation}
v_x(t_f) = 2v_0\sin\theta
\end{equation}

## Moto armonico

Il **moto armonico** è un moto oscillatorio di un punto $Q$ definito a partire da un punto $P$ che si muove di moto circolare uniforme su una circonferenza, considerando la sua proiezione $Q$ sul diametro della circonferenza (ovvero su un asse cartesiano).

\begin{figure}[ht]
\includegraphics[width=0.35\textwidth]{moto_armonico}
\centering
\caption{Moto parabolico}
\end{figure}

Dal moto circolare uniforme, si sa che la posizione della proiezione di $P$, ovvero $Q$, è data da:

\begin{equation} \label{eq:moto_armonico}
x(t) = R \cos(\omega t)
\end{equation}

Da cui si deriva

- Velocità:

\begin{equation}
v(t) = -R \omega \sin (\omega t)
\end{equation}

- Accelerazione:

\begin{equation}
a(t) = -R \omega^2 \cos (\omega t) = - \omega^2 x(t)
\end{equation}

$\omega$ è detta **pulsazione** del moto.

Se $Q$ parte da un punto diverso si individuato da un angolo $\phi$, detto **fase**, allora:

\begin{equation}
\begin{aligned}
& x(t) = R \cos(\omega t + \phi) \\
& v(t) = - R \sin(\omega t + \phi) \\
& a(t) = -R \sin(\omega t + \phi)
\end{aligned}
\end{equation}

# Dinamica del punto materiale

Prima legge di Newton (o legge dell'inerzia)
: Se un corpo non ha nessuna forza, manterrà la sua velocità"

Seconda legge di Newton
: Se un corpo ha accelerazione $a = [m/s^2]$ e massa $m = [kg]$ allora necessità di una forza

\begin{equation}
\vec F = m \vec a
\end{equation}

per raggiungerla. Dimensionalmente la forza $F$ è data da $[kg m / s^2]$ che si indicherà con $[N]$ (Newton).

Terza legge di Newton
: Dati due corpi $1$ e $2$ che esercitano reciprocamente delle forze, queste sono uguali e contrarie

\begin{equation} \label{eq:newton3}
F_{12} = - F_{21}
\end{equation}

Per capire l'esempio, Newton cita:

_«Ad ogni azione corrisponde una reazione pari e contraria. Se qualcuno spinge una pietra col dito, anche il suo dito viene spinto dalla pietra. Se un cavallo tira una pietra legata ad una fune, anche il cavallo è tirato ugualmente verso la pietra: infatti la fune distesa tra le due parti, per lo stesso tentativo di allentarsi, spingerà il cavallo verso la pietra e la pietra verso il cavallo; e di tanto impedirà l'avanzare dell'uno di quanto promuoverà l'avanzare dell'altro. Se un qualche corpo, urtando in un altro corpo, in qualche modo avrà mutato con la sua forza il moto dell'altro, a sua volta, a causa della forza contraria, subirà un medesimo mutamento del proprio moto in senso opposto. A queste azioni corrispondono uguali mutamenti, non di velocità, ma di moto. I mutamenti delle velocità, infatti, effettuati allo stesso modo in direzioni contrarie, in quanto i moti sono modificati in uguale misura, sono inversamente proporzionali ai corpi.»_

>**_TIP:_** tutte le forze che agiscono su un corpo sono forze da contatte, ad eccezione della forza di gravità.

Le leggi di Newton, in generale, introducono un importante concetto: la forza $F$ causa il moto, la massa $m$ di un corpo determina quanto il corpo si oppone a tale moto.

Forza risultante
: Se $F_1$ e $F_2$ formano un angolo $\theta$, allora la risultante è:

\begin{equation}
F = \sqrt{F_1^2 + F_2^2 + 2 F_1 F_2 \cos \theta}
\end{equation}

## Attrito radente

L'attrito radente si manifesta per oggetti solidi che strisciano su superfici solide. Vi sono due tipologie di attrito:

- Attrito statico: si oppone al movimento del corpo, produce una forza ugualmente contraria a quella applicata. Il massimo di forza applicate è dato da $F_s = u_s N$ dove $u_s$ è detto coefficiente d'attrito e $N$ è la forza normale.
- Attrito dinamico: l'oggetto inizia a muoversi quando viene applicata una forza maggiore a $F_s$. A questo punto la forza contraria dell'attritto cambia, e diventa $F_k = u_k N$ dove $u_k$ è detto coefficiente di attrito dinamico.

## Piano inclinato

Una massa $m$ è posta su un piano inclinato a un angolo $\theta$.

\begin{figure}[ht]
\includegraphics[width=0.35\textwidth]{piano_inclinato}
\centering
\caption{Piano inclinato} \label{fig:piano_inclinato}
\end{figure}

Per modellare il problema, gli assi $x$ e $y$ sono posti rispettivamente in maniera parallela e perpendicolare rispetto al piano. In questo modo la posizione sull'asse $y$ del corpo non varia, ma cambia soltanto quella su $x$.

Equazioni di base
: La forza normale del piano è associata alla componente verticale. Si danno le equazioni:

\begin{equation}
ma_x = mg\sin\theta
\label{eq:piano_x}
\end{equation}

\begin{equation}
\begin{aligned}
N = ma_y + mg\cos\theta \\
ma_y = N - mg\cos\theta
\end{aligned}
\label{eq:piano_y}
\end{equation}

poichè la componente verticale è bilanciata dalla forza peso, l'equazione \ref{eq:piano_y} diventa $ma_y = 0$. Nell'equazione \ref{eq:piano_x} si può semplificare la massa, ottenendo:

\begin{equation}
a_x = g\sin\theta
\end{equation}

Quindi il corpo scivolerà sul piano a un'accelerazione $g \sin\theta$.

Velocità in funzione della posizione
: Dall'equazione del moto $x(t) = \frac{1}{2} a_x t^2$ si ottiene la velocità $v_x(t) = a_x t$, dunque $t = \frac{v_x}{a_x}$, dunque $x = \frac{v_x^2}{2a_x}$, che permette di calcolare la velocità in funzione della posizione:

\begin{equation}
v_x = \sqrt{2xa_x} = \sqrt{2xg\sin\theta}
\end{equation}

Piano inclinato con attrito
: Aggiungendo un attrito al piano, l'equazione della componente verticale resta $N = mg\cos\theta$. Quella orizzontale diventa:

\begin{equation}
ma_x = mg\sin\theta - f
\end{equation}

$f$ varia, a seconda che sia attrito statico o dinamico. Sia $\theta^*$ il massimo angolo tale che il corpo non abbia accelerazione, ovvero l'angolo in cui si manifestra l'attrito statico, dato $F_s = u_s N$ con $N = mg\cos\theta$ si ottiene:

\begin{equation}
a_x = mg\sin\theta^* = u_s mg\cos\theta^*
\end{equation}

dividendo per $mg$:

\begin{equation}
\begin{aligned}
& \sin\theta^* = u_s \cos\theta^* \\
& u_s = \tan\theta^*
\end{aligned}
\end{equation} 

quando $\theta > \theta^*$ allora si manifesta l'attrito dinamico, e l'equazione diventa:

\begin{equation}
\begin{aligned}
& mg\sin\theta - u_k mg \cos\theta = ma_x \\
& a_x = g(\sin\theta - u_k \cos\theta)
\end{aligned}
\end{equation}

## Pendolo semplice

Una massa $m$ è appesa a una fune di lunghezza $L$ che oscilla attorno a un punto fisso.

\begin{figure}[ht]
\includegraphics[width=0.35\textwidth]{pendolo_semplice}
\centering
\caption{Pendolo semplice}
\end{figure}

Si sceglie come riferimento per l'asse $y$ l'asse coincidente con la direzione della fune, e come asse $x$ l'asse ortogonale a $y$

Sul punto agiscono la forza peso $P = mg$ verso il basso e la tensione della fune $T$. Sull'asse orizzontale appare solo la forza peso:

\begin{equation}
P_x = mg\sin\theta
\label{eq:forza_x}
\end{equation}

Analizzando il triangolo rettangolo che ha per ipotenusa la fune $L$ e per cateto orizzontale l'ampiezza $x$ allora:

\begin{equation}
x = L \sin\theta
\end{equation}

da cui:

\begin{equation}
\sin\theta = \frac{x}{L}
\end{equation}

riscrivendo la \ref{eq:forza_x}:

\begin{equation}
P_x = mg \frac{x}{L}
\end{equation}

applicando ils econdo principio della dinamica:

\begin{equation}
ma = -mg \frac{x}{L}
\end{equation}

ovvero:

\begin{equation}
a = -g \frac{x}{L}
\end{equation}

ricordando che $\omega = \sqrt{\frac{g}{L}}$:

\begin{equation}
a = -g \omega^2
\end{equation}

ricordando inoltre che $w = 2 \pi T$, allora si ottiene il periodo:

\begin{equation}
T = 2 \pi \sqrt{\frac{L}{g}}
\end{equation}

dunque per piccole oscillazioni il pendolo si comporta come un oscillatore armonico di periodo $T$.

## Gravitazione

Dati due corpi a una distanza $r$ l'equazione di gravitazione universale descrive l'attrazione tra i due:

$$
\vec F_{12} = \vec F{21} = G \frac{m_1 m_2}{r^2} \hat r
$$

$G$ è la costante di gravitazione universale.

$$
G = 6.67 \cdot 10^{-11} Nm^2kg^{-2}
$$

# Lavoro ed energia

## Lavoro

Si definisce lavoro elementare $W$ di una forza $\vec F$ associata a uno spostamento $d \vec s$ la forma differenziale data dal prodotto scalare tra $\vec F$ e $ds$:

\begin{equation}
dW = \vec F \cdot d \vec s
\end{equation}

L'unità di misura del lavoro è $[Nm] = [J]$ detta Joule.

Data una forza $F$ che cambia rispetto al tempo, il lavoro totale svolto tra $x_1$ e $x_2$ è l'integrale di $F$ tra $x_1$ e $x_2$:

\begin{equation}
W = \int_{x_1}^{x_2} F(x) dx
\end{equation}

Il lavoro si misura come $[L] = [FL] = [Nm]$ cioè Newton per metro, definito come Joule $J$.

## Energia cinetica

Si definisce energia cinetica la quantità:

\begin{equation}
K = \frac{1}{2}mv^2
\end{equation}

L'unità di misura dell'energia cinetica è il Joule.

Se un corpo di massa $m$ modifica la velocità da $v_i$ a $v_f$ allora varia l'energia cinetica:

\begin{equation}
\Delta K = K_f - K_i = \frac{1}{2} m v_f^2 - \frac{1}{2} m v_i^2
\end{equation}

## Teorema dell'energia cinetica

Il teorema del lavoro e dell'energia cinetica afferma che il lavoro svolto dalla forza risultante su un corpo è uguale alla differenza nell'energia cinetica, ovvero:

\begin{equation}
W = \Delta K
\label{eq:teorema_energia_cinetica}
\end{equation}

### Dimostrazione

Per dimostrarlo si ricorda che:

\begin{equation}
v^2 = v^2_0 + 2a(x - x_0)
\end{equation}

applicando il secondo principio della dinamica, sostituendo $a = F/m$:

\begin{equation}
v^2 = v^2_0 + 2\frac{F}{m}d
\end{equation}

con $d = (x_1 - x_2)$ distanza percorsa. Portando la forza da una parte e ciò che riguarda la particella dall'altra:

\begin{equation}
\frac{1}{2}mv^2_1 - \frac{1}{2}mv^2_2 = Fd
\label{eq:energia_1}
\end{equation}

sostituendo:

\begin{equation}
K_1 - K_2 = L
\end{equation}

## Energia potenziale

Una forza si dice conservatica se dipende solo dalla posizione nello spazio del punto materiale.

Si definisce variazione dell'energia potenziale $U$ il lavoro svolto dalla forza conservativa cambiata di segno:

\begin{equation}
\Delta U = -W
\label{eq:definizione_energia_potenziale}
\end{equation}

Sostituendo al lavoro la definizione generale per una forza variabile:

\begin{equation}
\Delta U = - \int_{x_1}^{x_2} F(x) dx
\end{equation}

## Energia meccanica

Si definisce energia meccanica $E$ la somma dell'energia potenziale $U$ e dell'energia cinetica $K$:

\begin{equation}
E = U + K
\end{equation}

## Principio di conservazione dell'energia meccanica

Il principio di conservazione dell'energia meccanica afferma che in un sistema in cui agiscono solo forze conservative allora l'energia meccanica si conserva, ovvero:

\begin{equation}
K_i + U_i = K_f + U_f
\end{equation}

Il significato è che dato un corpo sotto l'effetto di una forza $F(x)$ che accelera e decelera, vi è un valore:

\begin{equation}
E = \frac{1}{2}mv^2 + U(x)
\end{equation}

che non cambia mai. Ad esempio, nel caso della gravità $U = -mgy$, e quindi:

\begin{equation}
E_1 = \frac{1}{2}mv_1^2 + mgy_1 = \frac{1}{2}mv_2^2 + mgy_2 = E_2
\end{equation}

### Dimostrazione

Dal teorema dell'energia cinetica nella \ref{eq:teorema_energia_cinetica} e dalla definizione di energia potenziale nella \ref{eq:definizione_energia_potenziale}, si ottiene:

\begin{equation}
\Delta K = - \Delta U
\end{equation}

ovvero:

\begin{equation}
K_f - K_i = - U_f + U_i
\end{equation}

spostando i termini:

\begin{equation}
K_f + U_f = K_i + U_i
\end{equation}

# Dinamica traslazionale

## Sistemi di punti materiali

Un sistema di punti materiali è una collezione di $n > 1$ punti materiali. Questo tipo ha 6 variabili indipendenti: 3 per descrivere la posizione del punto in $(x, y, z)$ e tre per descriverne la rotazione attorno agli assi.

Sul sistema di punti materiai agiscono molte forze, interne ed esterne. Sulla particella $i$ la somma delle interazioni interne ed esterne è:

\begin{equation}
F_i = m_i \frac{d^2}{d^2t} \vec r_i
\end{equation}

la terza legge di Newton permette di eliminare le forze interne, in quanto uguali e contrarie, e considerare solo la somma delle forze esterne:

\begin{equation}
\vec F = M \frac{d^2}{d^2 t} \vec r_{CM}
\label{eq:traslatorio_1}
\end{equation}

con $M = \sum_i m_i$ massa totale del sistema e $\vec r_{CM}$ vettore del centro di massa, definito come media pesata dei vettori posizionali di ciascuno dei punti materiali del sistema:

\begin{equation}
\vec r_{CM} = \frac{\sum_i m_i \vec r_i}{M} = \frac{\sum_i m_i \vec r_i}{\sum_i m_i}
\end{equation}

## Potenza

Si definisce potenza il lavoro svolto in una certa unità di tempo, ovvero la derivata del lavoro rispetto al tempo:

\begin{equation}
P = \frac{dL}{dt}
\end{equation}

L'unità di misura è $[P] = [Js] = [W]$ detta Watt.

## Quantità di moto

Si definisce quantità di moto $\vec p$ il prodotto scalare tra la massa e la velocià:

\begin{equation}
\vec p = m \vec v
\end{equation}

La seconda legge di Newton si può riscrivere definendola come la derivata della quantità di moto rispetot al tempo:

\begin{equation}
\vec F = \frac{d \vec p}{dt}
\label{eq:definizione_forza_q}
\end{equation}

## Impulso e teorema dell'impulso

Si definisce impulso $I$ di una forza $F$ in un lasso di tempo $\Delta t = t_2 - t_1$ la grandezza:

\begin{equation}
\vec I = \int_{t_1}^{t_2} \vec F dt
\end{equation}

Il teorema dell'impulso afferma che l'impulso di una forza in un lasso di tempo è uguale alla variazione della quantità di moto in quel lasso di tempo:

\begin{equation}
\vec I = \Delta \vec p
\end{equation}

## Principio di conservazione della quantità di moto

Il principio di conservazione della quantità di moto afferma che se la forza risultante su un corpo è uguale a 0, allora la quantità di moto si conserva. Quindi:

\begin{equation}
\vec p_i = \vec p_f
\end{equation}

### Dimostrazione

Dalla nuova definizione di forza nella \ref{eq:definizione_forza_q} allora se $\vec F = 0$ anche la variazione della quantità di moto è nulla.

## Urti

Dati due corpi che si scontrano, vi sono due possibilità:

- Urti anelastici: l'energia cinetica non si conserva. In pratica, i due corpi continuano a muoversi assieme alla stessa velocità.
- Urti elastici: l'energia cinetica si conserva.

Si considerino due corpi di massa $m_1$ e $m_2$ con velocità $v_1$ e $v_2$ prima dell'urto e velocità $v_1'$ e $v_2'$ dopo l'urto. In entrambi i casi vale la legge di conservazione del momento, quindi:

\begin{equation}
m_1 v_1 + m_2 v_2 = m_1 v_1' + m_2 v_2'
\label{eq:urti_1}
\end{equation}

Urti anelastici
: Negli urti anelastici i due corpi, dopo l'urto, hanno una velocità comune $v' = v_1' = v_2'$, quindi:

\begin{equation}
m_1 v_1 + m_2 v_2 = v'(m_1 + m_2)
\end{equation}

Di conseguenza:

\begin{equation}
v' = \frac{m_1 v_1 + m_2 v_2}{m_1 + m_2}
\end{equation}

Urti elastici
: Negli urti elastici invece si conserva l'energia cinetica, quindi:

\begin{equation}
\frac{1}{2} m_1 v_1^2 + \frac{1}{2} m_2 v_2^2 = \frac{1}{2} m_1 (v_1')^2 + \frac{1}{2} m_2 (v_2')^2
\label{eq:urti_2}
\end{equation}

La \ref{eq:urti_2} assieme alla \ref{eq:urti_1} ha soluzioni:

\begin{equation}
\begin{aligned}
& v_1' = v_1 \frac{m_1 - m_2}{m_1 + m_2} + v_2 \frac{2 m_2}{m_1 + m_2} \\
& v_2' = v_2 \frac{m_2 - m_1}{m_1 + m_2} + v_1 \frac{2 m_1}{m_1 + m_2}
\end{aligned}
\end{equation}

# Dinamica rotazionale

Un sistema di punti oltre a traslare può ruotare attorno a un punto $O$ detto polo.

## Corpi rigidi

Un corpo rigido è un sistema di punti materiali in cui le distanze tra ogni coppia di punti non cambia. Quindi il centro di massa è in una posizione immutabile.

Un corpo rigido con infiniti punti materiali è trattato come un insieme di volumi $dV$, la cui somma da il volume del corpo:

\begin{equation}
V = \int_{V} dV
\end{equation}

ogni volumetto ha una massa $dm$, quindi la massa totale è:

\begin{equation}
M = \int dm
\end{equation}

ogni volumetto ha densità:

\begin{equation}
\rho = \frac{dm}{dV}
\end{equation}

quindi:

\begin{equation}
M = \int_V dV
\end{equation}

Nn corpo rigido si dice omogeneo se $\rho$ è costante. Vale:

\begin{equation}
M = \rho \int_V dV
\end{equation}

Il centro di massa dunque è:

\begin{equation}
\vec r_{CM} = \frac{\int_V \vec r \rho dV}{M}
\end{equation}

se è omogeneo:

\begin{equation}
\vec r_{CM} = \frac{1}{V} \int_V \vec r dV
\end{equation}

## Rotazione con un'accelerazione costante

Sia $P$ un punto in un corpo rigido distante $r$ dal centro a un angolo $\theta$ dall'asse $x$. Se il corpo sta ruotando allora $P$ si muove sulla circonferenza di raggio $r$. La distanza percorso in un certo spostamento di angolo $\theta$ è:

\begin{equation}
s = \theta r
\end{equation}

derivando:

\begin{equation}
\frac{d}{dt}s = r \frac{d}{dt} \theta
\end{equation}

si ricorda che la velocità angolare è la derivata dell'angolo percorso:

\begin{equation}
\omega = \frac{d}{dt} \theta
\end{equation}

quindi a sinistra appare la velocità tangenziale:

\begin{equation}
v_T = \frac{d}{dt} s = r \omega
\end{equation}

inoltre l'accelerazione angolare è:

\begin{equation}
\alpha = \frac{d}{dt} \omega = \frac{d^2}{dt^2} \theta
\end{equation}

similmente l'accelerazione centripeta è:

\begin{equation}
a_r = \frac{v^2}{r}
\end{equation}

utilizzando33 le equazioni del moto dalla cinematica:

\begin{equation}
\begin{aligned}
& \theta = \theta_0 + \omega_0 t + \frac{1}{2} \alpha t^2 \\
& \omega = \omega_0 + \alpha t \\
& \omega^2 = \omega_0^2 + 2 \alpha (\theta - \theta_0)
\end{aligned}
\end{equation}

Momento di inerzia
: Dato un punto materiale di massa $m$ che ruota s una circonferenza di raggio $r$ posto sull piano $O_{xy}$ intorno all'origine degli assi, si definisce momento di inerzia la tendenza del punto materiale ad opporsi alla rotazione ed è espresso dalla relazione:

\begin{equation}
I = m r^2
\end{equation}

## Momento torcente

Il momento torcente $\vec \tau$ è l'equivalente rotazionale delle forza, ed è definito come il prodotto vettoriale tra la posizione $\vec r$ su cui è applicata la forza $F$ (braccio della forza) e la forza stessa:

\begin{equation}
\vec \tau = \vec r \times \vec F
\end{equation}

Il modulo è $\tau = r F \sin(\theta)$ con $\theta$ angolo compreso tra $\vec r$ e $\vec F$; la direzione è perpendicolare al piano individuato da $\vec r$ e $\vec F$; il verso si ottiene con la regola della mano destra.

## Momento di inerzia

Il momento di inerzia $I$ è l'equivalente rotazionale della messa, cioè descrive quanto il corpo si oppone alla rotazione. Dato un punto che ruota a una distanza $r$ da un polo $O$ si definisce il momento di inerzia come il prodotto tra la sua massa $m$ ed $r$:

\begin{equation}
I = m r^2
\end{equation}

Inerzia in un sistema di punti materiali
: In un sistema di punti materiali l'inerzia è data dalla somma dell'inerzia dei punti.

Inerzia in un corpo rigido
: Il momento di inerzia in un corpo rigido è dato dalla somma del momento di inerzia dei volumetti, ovvero:

\begin{equation}
I = \int_{V} r^2 dm
\end{equation}

Teorema di Huygens-Steiner
: Il teorema di Huygens-Steiner enuncia:

\begin{equation}
I = I_c + md^2
\end{equation}

Dove $I_c$ è il momento di inerzia rispetto all'asse parallelo dato passante per il centro di massa, e $d^2$ è la distanza tra i due assi paralleli.

## Momento angolare

Il momento angolare $\vec L$ è l'equivalente rotazionale della quantità di moto. Si definisce come il prodotto vettoriale tra il vettore $\vec r$ che congiunge il polo con la posizione del punto materiale e la quantità di moto:

\begin{equation}
\vec L = \vec r \times \vec p
\label{eq:definizione_momento_inerzia}
\end{equation}

Il modulo è $L = r p \sin (\theta)$ con $\theta$ angolo compreso tra i vettori $\vec r$ e $\vec p$; la direzione è perpendicolare al piano individuato dai vettori $\vec r$ e $\vec p$; il verso si ottiene con la regola della mano destra.

Si osserva che derivando la \ref{eq:definizione_momento_inerzia} rispetto al tempo si ottiene:

\begin{equation}
\vec \tau = I \vec \alpha
\end{equation}

Infatti il momento torcente $\vec \tau$ lega il momento di inerzia (l'equivalente della massa) all'accelerazione $\alpha$ ed è ciò che causa un aumento del momento angolare.

Momento angolare in un sistema di punti materiali
: In un sistema di punti materiali il momento angolare è dato dalla somma del momento angolare di tutti i punti del sistema.

Momento angolare nei corpi rigidi
: Nei corpi rigido il momento angolare è definito sulla componente del momento angoalre lungo la direzione dell'asse di rotazione $\vec L_z$:

\begin{equation}
\vec L_z = I \vec \omega
\end{equation}

derivandola si ottiene la legge del moto rotatorio:

\begin{equation}
\frac{d \vec L}{dt} = - \vec v_O \times M \vec v_{CM} + \tau
\end{equation}

## Energia e lavoro nel moto rotazione

Il lavoro è definito come:

\begin{equation}
W = \int_{\theta_i}^{\theta_f} M d \theta
\end{equation}

L'energia cinetica è definita come:

\begin{equation}
E = \frac{1}{2} I_z \omega^2
\end{equation}

Il teorema dell'energia cinetica:

\begin{equation}
W = \Delta E = \frac{1}{2} I_z (\omega_f^2 - \omega_i^2)
\end{equation}

## Moto di puro rotolamento

Il moto di puro rotolamento combina un moto di traslazione dal centro di massa con velocità $v_{CM}$ e un moto rotatorio attorno al centro di massa con velocità angolare $\omega$. Poichè vi è rotoalemtno puro vale $\vec v_C = \vec v_{CM} + \vec \omega \times \vec r = 0$ dove $C$ è il punto di contatto, che è fermo rispetto al piano.

Le forza agenti sul corpo sono una forza esterna $F$, la forza di attrito $f$, il peso $mg$ con la relativa reazione normale $N$.

Valgono:

\begin{equation}
\begin{aligned}
& \text{Asse x}: & F - f = ma_{CM} \\
& \text{Asse y}: & N - mg = 0 \\
& \text{Relazione tra accelerazione lineare e angolare}: & a_{CM} = \alpha r \\
& \text{Momento torcente rispetto al centro di massa}: & \tau = I \alpha \\
& \text{Espressione per l'accelerazione angolare}: & \alpha = \frac{\tau}{I} \\
& \text{Relazione con l'accelerazione del centro di massa}: & a_{CM} = \alpha r \\
& \text{Somma delle forze esterne}: & rF = I \alpha \\
& \text{Momento torcente totale}: & a_{CM} = \frac{r^2 F}{I_{CM} + m R^2} \\
& \text{Accelerazione finale}: & a_{CM} = F \frac{1}{m(a + \frac{I_{CM}}{m R^2}} \\
& \text{Forza massima prima dello slittamento}: & F \leq \mu_s mg (1 + \frac{mR^2}{I})
\end{aligned}
\end{equation}

\appendix

# Vettori

Un vettore è caratterizzato da:

- Modulo: quanto è lungo
- Direzione: l'angolo del vettore relativo a un certo asse, solitamente l'asse $x$.
- Verso

## Versori

Si definiscono versori quei vettori che in uno spazio $\mathbb{R}^3$ puntano ai tre assi: $\hat i = (1, 0, 0), \hat j = (0, 1, 0), \hat k = (0, 0, 1)$. Ogni vettore si può scrivere come $\vec v = \hat i V_x + \hat j V_y + \hat k V_z$ dove $V_x, V_y, V_z$ sono detti componenti.

In un spazio a due dimensioni vale:

\begin{equation}
V_x = |V| \cos (\theta)
\end{equation}

\begin{equation}
V_y = |V| \sin (\theta)
\end{equation}

Facendo le inverse:

\begin{equation}
|V| = \sqrt{V_x^2 + V_y^2}
\end{equation}

\begin{equation}
\theta = \tan^{-1} \frac{V_x}{V_y}
\end{equation}

## Somma di vettori

Dati due vettori $\vec v, \vec u \in \mathbb{R}^k$, si calcola $\vec v + \vec u$ come segue:

- Date le componenti: siano $\vec v = (v_1, \cdots, v_k), \vec u = (u_1, \cdots, u_k)$ allora si sommano le singole componenti $\vec v + \vec u = (v_1 + u_1, \cdots, v_k + u_k)$.
- Dati i moduli e l'angolo $\theta$ tra i due vettori: si applica il teorema del coseno, e si calcola $|\vec v + \vec u| = \sqrt{|\vec v|^2 + |\vec u|^2 + |\vec v| \cdot |\vec u| \cos \theta}$

Per la somma di vettori valgono la proprietà commutativa, la proprietà associativa, l'esistenza dell'elemento neutro e l'esistenza dell'elemento opposto.

## Prodotto vettoriale

Dati due vettori $\vec v, \vec u \in \mathbb{R}^k$, l'operazione di moltiplicazione $\times : \mathbb{R}^3 \times \mathbb{R}^3 \rightarrow \mathbb{R^3}$ è calcolata per $\vec v \times \vec u$ come segue:

- Il modulo è $|\vec v \times \vec u| = |\vec v| \cdot |vec u| \cdot \sin \theta$
- La direzione è la direzione ortagonale che contiene i vettori $\vec v$ e $\vec u$

# Media MOG

Si calcola la media dei coefficienti $c_1$ e $c_2$:

$$
c = \frac{c_1 + c_2}{2}
$$

poi si calcola la media degli esponenti $e_1$ e $e_2$:

$$
e = \frac{e_1 + e_2}{2}
$$

si analizza la media dei coefficienti:

- se la somma è pari allora il risultato è $c^e$
- se la somma è dispari allora si pone $e = e - 0.5$ e il risultato è $3 \cdot c^e$