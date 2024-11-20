package com.example.foodapp.ui.shopping

import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.example.foodapp.R
import com.example.foodapp.adapter.HomeViewPagerAdapter
import com.example.foodapp.databinding.FragmentHomeBinding
import com.example.foodapp.databinding.FragmentLogInBinding
import com.google.android.material.tabs.TabLayoutMediator

class HomeFragment : Fragment(R.layout.fragment_home) {
    private lateinit var binding:FragmentHomeBinding
    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        binding = FragmentHomeBinding.inflate(layoutInflater)
        return binding.root
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        val adt = HomeViewPagerAdapter(childFragmentManager, lifecycle)
        binding.viewpagerHome.adapter = adt

        TabLayoutMediator(binding.tabLayout, binding.viewpagerHome){tab, position ->
            when (position) {
                0 -> tab.text= "All"
                1 -> tab.text= "Food"
                2 -> tab.text= "Drink"
            }
        }.attach()
    }

}