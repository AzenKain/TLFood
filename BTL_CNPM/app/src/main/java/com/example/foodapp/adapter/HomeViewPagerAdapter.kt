package com.example.foodapp.adapter

import androidx.fragment.app.Fragment
import androidx.fragment.app.FragmentManager
import androidx.lifecycle.Lifecycle
import androidx.viewpager2.adapter.FragmentStateAdapter
import com.example.foodapp.ui.categories.DrinkHomeFragment
import com.example.foodapp.ui.categories.FoodHomeFragment
import com.example.foodapp.ui.categories.MainHomeFragment

class HomeViewPagerAdapter(fragmentManager: FragmentManager, lifecycle: Lifecycle)
    : FragmentStateAdapter(fragmentManager, lifecycle) {
    override fun getItemCount(): Int {
        return 3
    }
    override fun createFragment(position: Int): Fragment {
        return when(position){
            0 -> {MainHomeFragment()}
            1 -> {FoodHomeFragment()}
            else -> {DrinkHomeFragment()}
        }
    }
}