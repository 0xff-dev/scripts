package com.example.uselistview;

import android.support.v7.app.ActionBar;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.ImageView;
import android.widget.ListView;
import android.content.Context;
import android.widget.TextView;

import java.util.ArrayList;
import java.util.List;

/**
 * learn ArrayAdapter, and use it
 * @author zhangshuang
 */
public class MainActivity extends AppCompatActivity {
    private List<Fruit> fruits = new ArrayList<Fruit>();

    @override
    protected void onCreate(Bundle saveInstanceState){
        super(saveInstanceState);
        this.init();
        FruitAdapter adapter = new FruitAdapter(MainActivity.this, R.layout.fruit_item, 
                                        this.fruits);
        Listview list = findViewById(R.id.list_view);
        list.setAdapter(adapter);
    }

    protected void init(){
        for(int i=0; i<5; i++){
            this.fruits.add(new Fruit("apple", R.drawable.apple));
        }
    }
}

class Fruit {
    private String name;
    private int imageId;

    public Fruit(String name, int imageId){
        this.name = name;
        this.imageId = imageId;
    }

    public String getName(){
        return this.name;
    }
    public int getId(){
        return this.imageId;
    }
}

class FruitAdapter extends ArrayAdapter<Fruit> {
    private int resourceId;
    public FruitAdapter(Context context, int textViewResourceId, List<Fruit> objects){
        super(context, textViewResourceId, objects);
        this.resourceId = textViewResourceId;  // 布局文件id
    }
    
    @override
    protected View getView(int position, View convertView, ViewGroup parent) {
        Fruit fruit = getitem(position);
        View view = LayoutInflater.from(getContext()).inflate(this.resourceId, parent, false);
        ImageView iv = (ImageView) view.findViewById(R.id.furit_image);
        TextView tv = (TextView) view.findViewById(R.id.fruit_name);
        iv.setImageResource(furit.getId());
        tv.setText(fruit.getName());
        return view;
    }
}