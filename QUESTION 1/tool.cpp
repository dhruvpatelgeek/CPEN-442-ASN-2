#include<bits/stdc++.h>
using namespace std;
#define x first
#define y second

/*
 * A string serarching algortithm
 *
 * input:  apple#applepenapple
 * Output: 000005000000050000
 * 
 * finds the occurance of the prefix of a string in that string in linear time
 */ 

vector<int> z_algorithm(string s){
	const int n=(int)s.length();
	vector<int> z(n); memset(&z[0],0x0,sizeof(int)*n);
	int x=0;int y=0;
	for(int i=1;i<n;i++){
		z[i]=(y<i)?0: min(y-i+1,z[i-x]);
		while(i+z[i]<n &&s[z[i]]==s[i+z[i]]){
			z[i]++;
		}
		if(i+z[i]-1>y){
			x=i;y=i+z[i]-1;
		}
	}
	return z;
}

void solve(){
	cout<<"\n enter the string : \n";
	string s; cin>>s;
	cout<<"\n the string is "<<s;
	vector<pair<char,int>> hist(26);
	map<char,int> histogram;
	for(int i=0;i<(int)s.size();i++){
		histogram[s[i]]++;
	}
	cout<<endl<<" histogram is \n";
	for(auto&&a:histogram){
		cout<<endl<<a.x<<" - "<<a.y;
	}
	int max_freq=0;
	char max_char='A';
	vector<char> int_map{'E','T','A','O','I','N','S','H','R','L','D','C','U','M','W','F','G','Y','P','B','V','K','J','X','Q','Z'};
	map<char,char> sub_mp;
	for(int i=0;i<26;i++){
		max_freq=-1;
		max_char='A';
		for(auto&&a:histogram){
			if(max_freq<a.y){
				max_freq=a.y;
				max_char=a.x;
			}
		}
		histogram.erase(max_char);
		sub_mp[max_char]=int_map[i];
	}
	cout<<endl<<" subtitution map is \n";
	for(auto&&a:sub_mp){
		cout<<endl<<a.x<<" -> "<<a.y;
	}
	for(int i=0;i<(int)s.size();i++){
		s[i]=sub_mp[s[i]];
	}
		cout<<"\n-------------------------\n";
		cout<<"\n the new string is "<<s;
		cout<<"\n-------------------------\n";
	string find;
	int choice=0;
	vector<int> zarr;
	int offset=0;
	while(1){
		cout<<"\n enter \n1 to find a string \n2 to subtitute\t";
		cin>>choice;
		if(choice==1){
			cin>>find;
			offset=find.size()+1;
			find+="#";
			find+=s;
			zarr=z_algorithm(s);
			cout<<endl;
			for(int i=0;i<s.size();i++){
				cout<<"{"<<zarr[i+offset]<<","<<s[i]<<"}";
			}
			cout<<endl;
			cout<<"\n SELECTED SUBSTRING ARE\n";
			for(int i=0;i<s.size();i++){
				if(zarr[i+offset]!=0)
					for(int j=0;j<offset&&i<(int)s.size();j++,i++)
						cout<<"{"<<zarr[i+offset]<<","<<s[i]<<"}";
			}
		}
		cout<<"\n the manunlly set subtition keys enter ther Key Value pairs K->V && V->K \t";
		char K,V;
		cin>>K>>V;
		if(K=='a'){
			break;
		}
		for(int i=0;i<(int)s.size();i++){
			if(s[i]==K){
				s[i]=V;
			} else if(s[i]==V) {
				s[i]=K;
			}
		}
		
		cout<<"\n-------------------------\n";
		cout<<"\n replaced "<<K<<" -> "<<V<<"\n";
		cout<<"\n the new string is "<<s;
	}	

}
int main(){
	solve();cout<<endl;
}
