package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)


// playfair creater from string
// playfair dycryptr
// calculate score

func playfair_genrate_matrix(keyword string)(map[string][2]int,[5][5]string){
	letters := [5][5]string{}
	var letter_string string;
	var alpha="abcdefghiklmnopqrstuvwxyz"
	var alpha_keyword =make(map[byte]int);
	alpha=keyword+alpha;
	for i:=0;i<len(alpha);i++{
		alpha_keyword[alpha[i]]=0
	}
	for i:=0;i<len(alpha);i++{
		if alpha_keyword[alpha[i]]==0{
			letter_string+=string(alpha[i])
		}
		alpha_keyword[alpha[i]]=1
	}
	//fmt.Println("ciper letter is ",letter_string,"with len",len(letter_string))

	for i,j,k:=0,0,0;i<5&&j<5&&k<25;{
		letters[j][i]= string(letter_string[k])
		i+=1
		k+=1
		if i%5==0{
			i=0
			j+=1
		}
	}
	var ans=make(map[string][2]int)
	var loc [2]int
	for i:=0;i<5;i++{
		//fmt.Println("")
		for j:=0;j<5;j++{
			loc[0]=i
			loc[1]=j
			ans[letters[i][j]]=loc
			//fmt.Print(letters[i][j]," ")
		}
	}
	//fmt.Println("")
	//for k,v:=range ans{
	//	fmt.Println(string(k),v)
	//}
	return ans,letters
}
func playfair_transform_string(letters string) [][2]byte{
	letters=strings.ToLower(letters)
	var diagraph [][2]byte
	var mono [2]byte
	if len(letters)%2==1{
		letters+="x"
	}
	for i:=0;i<len(letters);i+=2{
		mono[0]=letters[i];
		mono[1]=letters[i+1];
		if(mono[0]==mono[1]){
			letters="x"+letters
		}
		diagraph=append(diagraph, mono)
	}
	//for i:=0;i<len(diagraph);i++{
	//	fmt.Println(string(diagraph[i][0]),string(diagraph[i][1]));
	//}
	return diagraph
}

func calcualate_score(cipher string,words []string) int {
	cipher=strings.ToLower(cipher)
	var num_contained=0;
	for _, eachline := range words {
		if(strings.Contains(cipher,eachline)){
			//fmt.Println(num_contained," ",eachline)
			num_contained+= len(eachline)
		}
	}
	return num_contained;
}

func playfair_decrypt(table map[string][2]int,table_string [5][5]string,diagraph [][2]byte)string {
	var ans_string string
	var loc1,loc2 [2]int
	var a,b int
	for i:=0;i<len(diagraph);i++{
		loc1=table[string(diagraph[i][0])];
		loc2=table[string(diagraph[i][1])];
		if loc1[0]==loc2[0] {
			//row
			// assume l1 is further than l2
			if(loc1[1]<loc2[1]){
				a = loc1[1] - 1;
				if (a < 0) {
					a = 4
				}
				b = loc2[1] - 1;
				if (b < 0) {
					b = 4
				}

				ans_string += table_string[loc1[0]][a];
				ans_string += table_string[loc2[0]][b];
			} else {
				a = loc1[1] - 1;
				if (a < 0) {
					a = 4
				}
				b = loc2[1] - 1;
				if (b < 0) {
					b = 4
				}
				ans_string += table_string[loc1[0]][a];
				ans_string += table_string[loc2[0]][b];
			}
		} else if loc1[1]==loc2[1]{
			//column
			if(loc1[0]<loc2[0]){
				a=loc1[0]-1;
				if(a<0){
					a=4
				}
				b=loc2[0]-1;
				if(b<0){
					b=4
				}
				ans_string+=table_string[a][loc1[1]];
				ans_string+=table_string[b][loc2[1]];
			} else {
				a = loc1[0] - 1;
				if (a < 0) {
					a = 4
				}
				b = loc2[0] - 1;
				if (b < 0) {
					b = 4
				}
				ans_string += table_string[a][loc1[1]];
				ans_string += table_string[b][loc2[1]];
			}
		} else {
			//rectangle
			// loc 2 on the right loc 1 on the left

			if(loc1[1]>loc2[1]){
				ans_string+=table_string[loc1[0]][loc2[1]]
				ans_string+=table_string[loc2[0]][loc1[1]]
			} else {
				ans_string+=table_string[loc1[0]][loc2[1]]
				ans_string+=table_string[loc2[0]][loc1[1]]
			}


		}
	}
	//fmt.Println("ANS >> ",ans_string)
	return ans_string
}

var NUM_WORKERS=1;
var wg sync.WaitGroup
func main() {

	fmt.Println("Starting HILL WORKER CRACKER WITH [",NUM_WORKERS,"] WORKERS")
	file, err := os.Open("./resources/words_alpha.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var englishDict []string

	for scanner.Scan() {
		englishDict = append(englishDict, scanner.Text())
	}
	file.Close()

	file, err = os.Open("./resources/candidate_keys.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var candidate []string

	for scanner.Scan() {
		candidate = append(candidate, scanner.Text())
	}
	file.Close()

	var cipher_text string="GNCNTNIEVUPTOTCHZSWQTQFNMFMDTLEFXUPTSDFUEYTFQNYEMZGYTFYXPXNRZMECFITRWEYXXTNPNSADZETFYFSTZEDRFUTRWEYXNKZMUMOTFTYFMZQKNRPTHUTADLWMWZSOLFPHKOGYOTTQOHTRZMUDQMSYTRWENFUVPQDCNDOTCHETISYGHIZVZPGUTFPTHUACTQGYAVFGHIELPZEUOTAMUIMFXMUHISSZOTAMFTSUTRWEYGHIFVFTONMYRPXNTQLWMZENWEXNPKLYSHQSWDWEYXBTQIVLNPTPPKLYSHOTENQNYMHIFVUDTLGYTFHOWEHYXYAVEHTLEFTXSZPQPTHUOTTQLVNUTQFNMFPBPTZPOTUNSRHIEHTLGYSUSITRWENFUVOTFTYFMZQKNRQSWDWEYHTCZRGQTFPTOHHQVUPBPTHUPTSOYFMZGNSRHIFHTLMPNODCDTSOHIHTLUTBZFVUGYAVAHNOMWYXNUTQWKCHTVNOSNKMSQTRWEYXPBPTSNOHADFMEULWLYTYSOHIQNITLXADZFWOUMOTFDWZYHNCTQHNMWZFHEUXFTQGUMNOOTFTYFMZQKNRMASOHNWPZPKAIEAIXFSFONVHBCGQATBTYXMDHZMDTLRKNLONLOEXNFUVYXPBPTZPOTNAFNTCZRGQTFMDSZYHMDNLNDOTLMMZENGYAFMPAYHSNOLNZFBUTRWEFNHTLTEXTFATVHKUTRWEDCNEZHDLAIEFXUPTNEAVEHTLNFEUQSWDWENFUVOTFTYFMZQKNRQSWDWEHTATADUMNFTLZWNETKQINLHOTLHUXANPTPWETXNONSOTATADDEPZHYHWOTFTHZVUZMGYOTOTLMMZATGNCHMFKOYFEUONOTUDYHRNTPWEXYTXNFUVPQNLNBTFSOZSWQTQPTSEHZZHKRUMNOHNOTUDYHRNTPWESAHZZHMAVTUGYXBUEWUDNOTLMWZFNDOTZPTAZPTRPTOTUMUVSNSMSNAWGNUDTLGOAMTQEVLCTXPQFEHPZEIAEXQNWZTLEXQSWDWEDCOTLTPRUMUHRNIEPVUDTLZHDLAIEFXUPTSEPZQSWDWETLNFUVONVHKTZECSTRWECLUHCHOTLTZFBUTRWEYLKTQMLUHTFEAHNODFLHDLXYTLZWSUTRWEHCSYZFTLEXNLHSNTZFMOSZUHHNAXYXMDTLMZZSMZYFXWSNSEADZEZSWRUMNODCHULCHUHOTLZHDLAIEFXUCHNLAVFVLMNLMPZSWZTLZWSUTRWEZSWZAVAVEFVUFTZWUDTLKTZEDLNBNFUVONDFGCTXAITLZWDUPBPTSDTDGNATOTAMNPTPWETASOZMUMOTLMMZENOTZPTAZPTRTLNFAXLUOTZMOTLTZFXTPZOTZEEXCITRWEZSZPLOEXNEMFSOTLHTZXZTLFRTHODUWPGQVFFDZFHOPTHUNMFNTAZWAIFNVCWLNFTAEZVOYXSUTRWEYXMDTLMWADATOTZPTATFUEUMNLMPAIHOZPHWATTQLVONTNEVUHNKLULVDMNOOTFDZFHONLXETLEXIHMFOTENSZZWVHSNUFZPBSTRWENSAIOTZMAMZMZLHGADAUFTYHUHBCFYUMNOOTZMAMZMZLHGZENSWLWRHKZPGYHWTQONLCOHPSXUQSWDWEYXVUNEFHNTEXLCDLZHHWQSWDWEHFUGEZKHPTNDOTUIFSTRWERNHZAINLQGQSDLMZMUDLWQXZMOSZUHQSWDWEYXXERTMZPMZFBEIEZHSRYHMRNSHOZMQSIEXTNODFUGLWYHUMPTHUNMFNZSMOYTENUDTLHUXANPTPWESTEXSNTLZWNENDOTENRTEUSOQSWDWEZSGQNMTQEVNLDLDFENZHTMYXBEFNSEZPEDZMZPMPNODCSDLUHWNSAIHUSYTRWEOTAMNPTPWEXAFNSNNENEAVEFZKLYBUTRWEFEOTTQLVTLAYTKQITVCHZXRPYHDUPBPTHUEWPQQSWDWEOTLMMZENYXMDZMGQVFFUTRWEOBSRUDTLDFFNSUTRWEYHLXHQUFOTENAYHNTAMWZFOMNZMPNOGNNZMZQNQITRWEYXXMNIPTQNCTADFMEUYLWERNTPWESYTRWEYXMEEXZTESYHPTNDOTZMZLPZLUNLDLSNTLZWSUTRWEYXBUYXSNHEHPUMNOQVWZNWQNTDHRHGUEZFTLZWNDOTENUGYNDFFNNOTLWRATZWUMNOGNCNTNIEVUPTOTCHMUWETXUMOTFDZFHOPTSDFUEYTFYLBFUHQSWDWEYXXUTQFENSFUYXMWYXDEQWZUTUHOQSWDWEYXNKZMUMOTAMNPTPWETXNOPQWLDLFENRMWNREFNDOTCHQSWDWEDCOHLYBUTRWEMWCELECUACTQYHMUUMMUHCMZLNZFMOQSWDWEQNTECTDUZWHYDMNOTNAYBFUMOTAMNPTPWETXNOTLZVPTZPOTUNSRHIEHTLGYSFUMGYSUTNESTLMAHOQSWDWEMZVZULMDTLDFFNDMNOGNATOTAMUHISCHMUEXQSWDWEQSEUONEUQSWDWEYXVUGYAVEHTLHWHGAUTQOTFDHLMLHOYXBMYHMUHOYXMENHHEHPYHLVZFMOYHFEEHTLEFTXSZPQ"
	for i:=0;i<NUM_WORKERS;i++{
		wg.Add(1)
		hill_worker(cipher_text,englishDict,1,candidate[i]);
	}
	wg.Wait()
}

func next(key string) string{
	var new_key_byte [25]byte
	var new_key string
	for i:=0;i<25&&i<len(key);i++{
		new_key_byte[i]=key[i]
	}

	new_key_byte[24]=new_key_byte[24]+1
	if(new_key_byte[24]==123){
		new_key_byte[24]=97;
		for i:=23;i>=0;i--{
			new_key_byte[i]=new_key_byte[i]+1
			if(new_key_byte[i]==123){
				new_key_byte[i]=97
				continue;
			} else {
				break;
			}
		}
	}

	for i:=0;i<25;i++{
		new_key+=string(new_key_byte[i])
	}
	return new_key
}
func prev(key string) string{
	var new_key_byte [25]byte
	var new_key string
	for i:=0;i<25&&i<len(key);i++{
		new_key_byte[i]=key[i]
	}

	new_key_byte[24]=new_key_byte[24]-1
	if(new_key_byte[24]==96){
		new_key_byte[24]=122;
		for i:=23;i>=0;i--{
			new_key_byte[i]=new_key_byte[i]-1
			if(new_key_byte[i]==96){
				new_key_byte[i]=122
				continue;
			} else {
				break;
			}
		}
	}

	for i:=0;i<25;i++{
		new_key+=string(new_key_byte[i])
	}
	return new_key
}
func hill_worker(cipher_text string,englishDict []string, wrk_num int,starting_string string){
	var file_name = "WRKER_RESULT#"
	file_name+=strconv.Itoa(wrk_num)
	file_name+=".log"
	f, err := os.Create(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()



	defer wg.Done()
	var itr=0;
	var decrypted_text string
	var prev_best_decrypted_text string
	var prev_best_score int =0
	var prev_best_string string=starting_string
	var forward=100;
	var backwards=100;
	var direction bool =true
	for {
		itr++;
		if(itr%100==0){
			fmt.Println("WORKER #",wrk_num,"COMPLETED ",itr,"CHECKS");
		}
		table, table_string := playfair_genrate_matrix(starting_string);
		diagraph := playfair_transform_string(cipher_text)
		decrypted_text = playfair_decrypt(table, table_string, diagraph)
		curr_score := calcualate_score(decrypted_text, englishDict)
		if(curr_score<prev_best_score){
			if(direction){
				forward--;
				backwards=100;
			} else {
				backwards--;
			}
			if(forward==0){
				direction=false;
			} else if(backwards==0) {
				direction=true;
			}
			if(forward==0&&backwards==0){
				break;
			}
		} else {
			prev_best_score=curr_score
			prev_best_string=starting_string
			prev_best_decrypted_text=decrypted_text
		}
		if(direction){
			starting_string=next(starting_string);
		} else {
			starting_string=prev(starting_string);
		}
	}
	fmt.Println("WORKER EXIT",wrk_num,"max_score",prev_best_score);
	if(prev_best_score>4000){
		fmt.Println(prev_best_decrypted_text,"\n -------\n with KEY ",prev_best_string)
	}
	_, err2 := f.WriteString(prev_best_decrypted_text+"\n -------\n"+"with KEY "+prev_best_string)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}